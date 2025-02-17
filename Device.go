package onvif

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/jjbarbosa7/onvif/device"
	"github.com/jjbarbosa7/onvif/gosoap"
	"github.com/jjbarbosa7/onvif/networking"
	wsdiscovery "github.com/jjbarbosa7/onvif/ws-discovery"
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

// DeviceInfo struct contains general information about ONVIF device
type DeviceInfo struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareId      string
}

// Device for a new device of onvif and DeviceInfo
// struct represents an abstract ONVIF device.
// It contains methods, which helps to communicate with ONVIF device
type Device struct {
	params    DeviceParams
	endpoints map[string]string
	info      DeviceInfo
}

type DeviceParams struct {
	Xaddr      string
	Username   string
	Password   string
	HttpClient *http.Client
}

// GetServices return available endpoints
func (dev *Device) GetServices() map[string]string {
	return dev.endpoints
}

// GetDeviceInfo return available endpoints
func (dev *Device) GetDeviceInfo() DeviceInfo {
	return dev.info
}

// GetDeviceParams return available endpoints
func (dev *Device) GetDeviceParams() DeviceParams {
	return dev.params
}

// GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) ([]Device, error) {
	// Call a ws-discovery Probe Message to Discover NVT type Devices
	devices, err := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		return nil, err
	}

	nvtDevicesSeen := make(map[string]bool)
	nvtDevices := make([]Device, 0)

	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil, err
		}

		for _, xaddr := range doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs") {
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			if !nvtDevicesSeen[xaddr] {
				dev, _, err := NewDevice(DeviceParams{Xaddr: strings.Split(xaddr, " ")[0]})
				if err != nil {
					// TODO(jfsmig) print a warning
				} else {
					nvtDevicesSeen[xaddr] = true
					nvtDevices = append(nvtDevices, *dev)
				}
			}
		}
	}

	return nvtDevices, nil
}

func (dev *Device) getTimeDiff(resp *http.Response) (time.Duration, error) {
	doc := etree.NewDocument()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	resp.Body.Close()

	if err := doc.ReadFromBytes(data); err != nil {
		return 0, err
	}

	utcE1 := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/UTCDateTime")
	if utcE1 == nil {
		return 0, fmt.Errorf("UTCDateTime not found")
	}

	// Find the <Time> element
	timeEl := utcE1.FindElement("./Time")
	if timeEl == nil {
		return 0, fmt.Errorf("time element not found in UTCDateTime")
	}
	hourStr := strings.TrimSpace(timeEl.FindElement("Hour").Text())
	minStr := strings.TrimSpace(timeEl.FindElement("Minute").Text())
	secStr := strings.TrimSpace(timeEl.FindElement("Second").Text())

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return 0, err
	}
	min, err := strconv.Atoi(minStr)
	if err != nil {
		return 0, err
	}
	sec, err := strconv.Atoi(secStr)
	if err != nil {
		return 0, err
	}

	// Find the <Date> element
	dateEl := utcE1.FindElement("./Date")
	if dateEl == nil {
		return 0, fmt.Errorf("date element not found in UTCDateTime")
	}
	yearStr := strings.TrimSpace(dateEl.FindElement("Year").Text())
	monthStr := strings.TrimSpace(dateEl.FindElement("Month").Text())
	dayStr := strings.TrimSpace(dateEl.FindElement("Day").Text())

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, err
	}
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0, err
	}
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, err
	}

	utcDateTime := time.Date(year, time.Month(monthInt), day, hour, min, sec, 0, time.UTC)
	timeDiff := time.Now().UTC().Sub(utcDateTime)

	return timeDiff, nil
}

func (dev *Device) getSupportedServices(resp *http.Response) error {
	doc := etree.NewDocument()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body.Close()

	if err := doc.ReadFromBytes(data); err != nil {
		return err
	}

	services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
	for _, j := range services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}

	extension_services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Extension/*/XAddr")
	for _, j := range extension_services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}

	return nil
}

// NewDevice function construct a ONVIF Device entity
func NewDevice(params DeviceParams) (*Device, time.Duration, error) {
	dev := new(Device)
	dev.params = params
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+dev.params.Xaddr+"/onvif/device_service")

	if dev.params.HttpClient == nil {
		dev.params.HttpClient = new(http.Client)
	}

	// Attempt to get the time difference between the client and the camera; this is used for WS-Security
	// First try to get the time without authentication, as with some cameras (Axis), the authentication will fail
	// if the time difference between the client and the camera is too large.
	getTime := device.GetSystemDateAndTime{}
	resp, err := dev.CallMethod(getTime, time.Duration(0), true)
	if err != nil || resp.StatusCode != http.StatusOK {
		resp, err = dev.CallMethod(getTime, time.Duration(0), false)
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, time.Duration(0), fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %v", dev.params.Xaddr, err)
		}
	}
	timeDiff, err := dev.getTimeDiff(resp)
	if err != nil {
		return nil, time.Duration(0), fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %v", dev.params.Xaddr, err)
	}

	getCapabilities := device.GetCapabilities{Category: "All"}
	resp, err = dev.CallMethod(getCapabilities, timeDiff, false)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, timeDiff, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetCapabilities): %v", dev.params.Xaddr, err)
	}

	err = dev.getSupportedServices(resp)
	if err != nil {
		return nil, timeDiff, err
	}

	return dev, timeDiff, nil
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
func (dev Device) CallMethod(method interface{}, timeDiff time.Duration, omitSecurityHeader bool) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, err
	}
	return dev.callMethodDo(endpoint, method, timeDiff, omitSecurityHeader)
}

// CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method interface{}, timeDiff time.Duration, omitSecurityHeader bool) (*http.Response, error) {
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

	// Get method name
	methodName := reflect.TypeOf(method).Name()

	fmt.Printf("Calling %s on %s\n", methodName, endpoint)

	//Auth Handling
	if !omitSecurityHeader && dev.params.Username != "" && dev.params.Password != "" {
		soap.AddWSSecurity(dev.params.Username, dev.params.Password, timeDiff)
	}

	return networking.SendSoap(dev.params.HttpClient, endpoint, soap.String())
}
