package onvif

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/gofrs/uuid"
	"github.com/jjbarbosa7/onvif/device"
	"github.com/jjbarbosa7/onvif/gosoap"
	"github.com/juju/errors"
	"golang.org/x/net/ipv4"
)

// Xlmns XML Scheam
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tev":     "http://www.onvif.org/ver10/events/wsdl",
	"tptz":    "http://www.onvif.org/ver20/ptz/wsdl",
	"timg":    "http://www.onvif.org/ver20/imaging/wsdl",
	"tan":     "http://www.onvif.org/ver20/analytics/wsdl",
	"xmime":   "http://www.w3.org/2005/05/xmlmime",
	"wsnt":    "http://docs.oasis-open.org/wsn/b-2",
	"xop":     "http://www.w3.org/2004/08/xop/include",
	"wsa":     "http://www.w3.org/2005/08/addressing",
	"wstop":   "http://docs.oasis-open.org/wsn/t-1",
	"wsntw":   "http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw": "http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":    "http://www.w3.org/2006/05/addressing/wsdl",
}

// DeviceType alias for int
type DeviceType int

// Onvif Device Tyoe
const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

const (
	MULTI_RESPONSE_READ_BUFFER = 8192
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

// Device for a new device of onvif and DeviceInfo
// struct represents an abstract ONVIF device.
// It contains methods, which helps to communicate with ONVIF device
type Device struct {
	params    DeviceParams
	endpoints map[string]string
}

type DeviceParams struct {
	Xaddr    string
	Username string
	Password string
}

// GetServices return available endpoints
func (dev *Device) GetServices() map[string]string {
	return dev.endpoints
}

// GetDeviceParams return available endpoints
func (dev *Device) GetDeviceParams() DeviceParams {
	return dev.params
}

func (dev *Device) getTimeDiff(resp *http.Response) (time.Duration, error) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()

	systemDateTime, err := dev.DecodeSystemDateTime(data)
	if err != nil {
		return 0, err
	}

	if systemDateTime.UTCDateTime.IsZero() {
		return 0, fmt.Errorf("UTCDateTime not found")
	}

	timeDiff := time.Now().UTC().Sub(systemDateTime.UTCDateTime)

	return timeDiff, nil
}

func (dev *Device) getSupportedServices(resp *http.Response) (*Capabilities, error) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	capabiities, err := dev.DecodeCapabilities(data)
	if err != nil {
		return nil, err
	}
	if capabiities.Analytics.XAddr != "" {
		dev.addEndpoint("Analytics", capabiities.Analytics.XAddr)
	}
	if capabiities.Device.XAddr != "" {
		dev.addEndpoint("Device", capabiities.Device.XAddr)
	}
	if capabiities.Events.XAddr != "" {
		dev.addEndpoint("Events", capabiities.Events.XAddr)
	}
	if capabiities.Imaging.XAddr != "" {
		dev.addEndpoint("Imaging", capabiities.Imaging.XAddr)
	}
	if capabiities.Media.XAddr != "" {
		dev.addEndpoint("Media", capabiities.Media.XAddr)
	}
	if capabiities.PTZ.XAddr != "" {
		dev.addEndpoint("PTZ", capabiities.PTZ.XAddr)
	}

	return capabiities, nil
}

// NewDevice function construct a ONVIF Device entity
func NewDevice(params DeviceParams, httpClient *http.Client, omitSecurityHeader bool) (*Device, time.Duration, *Capabilities, error) {
	dev := new(Device)
	dev.params = params
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+dev.params.Xaddr+"/onvif/device_service")

	// Attempt to get the time difference between the client and the camera; this is used for WS-Security
	// First try to get the time without authentication, as with some cameras (Axis), the authentication will fail
	// if the time difference between the client and the camera is too large.
	getTime := device.GetSystemDateAndTime{}
	resp, err := dev.CallMethod(getTime, time.Duration(0), httpClient, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		if !omitSecurityHeader {
			resp, err = dev.CallMethod(getTime, time.Duration(0), httpClient, false)
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, time.Duration(0), nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %s", dev.params.Xaddr, err)
		}
	}
	timeDiff, err := dev.getTimeDiff(resp)
	if err != nil {
		return nil, time.Duration(0), nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %s", dev.params.Xaddr, err)
	}

	getCapabilities := device.GetCapabilities{Category: "All"}
	resp, err = dev.CallMethod(getCapabilities, timeDiff, httpClient, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		if !omitSecurityHeader {
			resp, err = dev.CallMethod(getCapabilities, timeDiff, httpClient, false)
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, timeDiff, nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetCapabilities): %s", dev.params.Xaddr, err)
		}
	}

	capabilities, err := dev.getSupportedServices(resp)
	if err != nil {
		return nil, timeDiff, nil, err
	}

	return dev, timeDiff, capabilities, nil
}

func (dev *Device) addEndpoint(Key, Value string) {
	//use lowCaseKey
	//make key having ability to handle Mixed Case for Different vendor devcie (e.g. Events EVENTS, events)
	lowCaseKey := strings.ToLower(Key)

	// Replace host with host from device params.
	if u, err := url.Parse(Value); err == nil {
		u.Host = dev.params.Xaddr
		Value = u.String()
	}

	dev.endpoints[lowCaseKey] = Value
}

// GetEndpoint returns specific ONVIF service endpoint address
func (dev *Device) GetEndpoint(name string) string {
	return dev.endpoints[name]
}

func (dev Device) buildMethodSOAP(msg string) (gosoap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg); err != nil {
		//log.Println("Got error")

		return "", err
	}
	element := doc.Root()

	soap := gosoap.NewEmptySOAP()
	soap.AddBodyContent(element)

	return soap, nil
}

// getEndpoint functions get the target service endpoint in a better way
func (dev Device) getEndpoint(endpoint string) (string, error) {
	// common condition, endpointMark in map we use this.
	if endpointURL, bFound := dev.endpoints[endpoint]; bFound {
		return endpointURL, nil
	}

	//but ,if we have endpoint like event、analytic
	//and sametime the Targetkey like : events、analytics
	//we use fuzzy way to find the best match url
	var endpointURL string
	for targetKey := range dev.endpoints {
		if strings.Contains(targetKey, endpoint) {
			endpointURL = dev.endpoints[targetKey]
			return endpointURL, nil
		}
	}
	return endpointURL, fmt.Errorf("target endpoint service not found")
}

// CallMethod functions call an method, defined <method> struct.
// You should use Authenticate method to call authorized requests.
func (dev Device) CallMethod(method interface{}, timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, err
	}
	return dev.callMethodDo(endpoint, method, timeDiff, httpClient, omitSecurityHeader)
}

// CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method interface{}, timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) (*http.Response, error) {
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		return nil, err
	}

	soap, err := dev.buildMethodSOAP(string(output))
	if err != nil {
		return nil, err
	}

	soap.AddRootNamespaces(Xlmns)
	soap.AddAction()

	//Auth Handling
	if !omitSecurityHeader {
		soap.AddWSSecurity(dev.params.Username, dev.params.Password, timeDiff)
	}

	return sendSoap(httpClient, endpoint, soap.String())
}

// sendSoap sends an ONVIF SOAP request and retries if 401 Unauthorized is received
func sendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(message))
	if err != nil {
		return nil, errors.Annotate(err, "Failed to create request")
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("SOAPAction", `""`)

	// First request attempt
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Failed to send request")
	}

	// If response is 401, retry with authentication
	if resp.StatusCode == http.StatusUnauthorized {
		// Close first response before retrying
		resp.Body.Close()

		// Resend request
		resp, err = httpClient.Do(req)
		if err != nil {
			return nil, errors.Annotate(err, "Failed to send authenticated request")
		}
	}

	return resp, nil
}

func sendProbe(interfaceName string, scopes, types []string, namespaces map[string]string) ([]string, error) {
	// Creating UUID Version 4
	uuidV4 := uuid.Must(uuid.NewV4())

	probeSOAP := buildProbeMessage(uuidV4.String(), scopes, types, namespaces)
	//probeSOAP = `<?xml version="1.0" encoding="UTF-8"?>
	//<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing">
	//<Header>
	//<a:Action mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</a:Action>
	//<a:MessageID>uuid:78a2ed98-bc1f-4b08-9668-094fcba81e35</a:MessageID><a:ReplyTo>
	//<a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
	//</a:ReplyTo><a:To mustUnderstand="1">urn:schemas-xmlsoap-org:ws:2005:04:discovery</a:To>
	//</Header>
	//<Body><Probe xmlns="http://schemas.xmlsoap.org/ws/2005/04/discovery">
	//<d:Types xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:dp0="http://www.onvif.org/ver10/network/wsdl">dp0:NetworkVideoTransmitter</d:Types>
	//</Probe>
	//</Body>
	//</Envelope>`

	return sendUDPMulticast(probeSOAP.String(), interfaceName)
}

func sendUDPMulticast(msg string, interfaceName string) ([]string, error) {
	c, err := net.ListenPacket("udp4", "0.0.0.0:0")
	if err != nil {
		return nil, err
	}
	defer c.Close()

	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return nil, err
	}

	p := ipv4.NewPacketConn(c)
	group := net.IPv4(239, 255, 255, 250)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		return nil, err
	}

	dst := &net.UDPAddr{IP: group, Port: 3702}
	data := []byte(msg)
	for _, ifi := range []*net.Interface{iface} {
		if err := p.SetMulticastInterface(ifi); err != nil {
			return nil, err
		}
		p.SetMulticastTTL(2)
		if _, err := p.WriteTo(data, nil, dst); err != nil {
			return nil, err
		}
	}

	if err := p.SetReadDeadline(time.Now().Add(time.Second * 1)); err != nil {
		return nil, err
	}

	var result []string
	for {
		b := make([]byte, MULTI_RESPONSE_READ_BUFFER)
		n, _, _, err := p.ReadFrom(b)
		if err != nil {
			if !errors.Is(err, os.ErrDeadlineExceeded) {
				return nil, err
			}
			break
		}
		result = append(result, string(b[0:n]))
	}
	return result, nil
}

func buildProbeMessage(uuidV4 string, scopes, types []string, nmsp map[string]string) gosoap.SoapMessage {
	namespaces := make(map[string]string)
	namespaces["a"] = "http://schemas.xmlsoap.org/ws/2004/08/addressing"

	probeMessage := gosoap.NewEmptySOAP()

	probeMessage.AddRootNamespaces(namespaces)

	var headerContent []*etree.Element

	action := etree.NewElement("a:Action")
	action.SetText("http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe")
	action.CreateAttr("mustUnderstand", "1")

	msgID := etree.NewElement("a:MessageID")
	msgID.SetText("uuid:" + uuidV4)

	replyTo := etree.NewElement("a:ReplyTo")
	replyTo.CreateElement("a:Address").SetText("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous")

	to := etree.NewElement("a:To")
	to.SetText("urn:schemas-xmlsoap-org:ws:2005:04:discovery")
	to.CreateAttr("mustUnderstand", "1")

	headerContent = append(headerContent, action, msgID, replyTo, to)
	probeMessage.AddHeaderContents(headerContent)

	//Содержимое Body
	probe := etree.NewElement("Probe")
	probe.CreateAttr("xmlns", "http://schemas.xmlsoap.org/ws/2005/04/discovery")

	if len(types) != 0 {
		typesTag := etree.NewElement("d:Types")
		if len(nmsp) != 0 {
			for key, value := range nmsp {
				typesTag.CreateAttr("xmlns:"+key, value)
			}
		}
		typesTag.CreateAttr("xmlns:d", "http://schemas.xmlsoap.org/ws/2005/04/discovery")
		var typesString string
		for _, j := range types {
			typesString += j
			typesString += " "
		}

		typesTag.SetText(strings.TrimSpace(typesString))

		probe.AddChild(typesTag)
	}

	if len(scopes) != 0 {
		scopesTag := etree.NewElement("d:Scopes")
		var scopesString string
		for _, j := range scopes {
			scopesString += j
			scopesString += " "
		}
		scopesTag.SetText(strings.TrimSpace(scopesString))

		probe.AddChild(scopesTag)
	}

	probeMessage.AddBodyContent(probe)

	return probeMessage
}
