package onvif

import (
	"bytes"
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
	"github.com/jjbarbosa7/onvif/xsd_onvif"
	"github.com/juju/errors"
)

// Xlmns XML Scheam
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tr2":     "http://www.onvif.org/ver20/media/wsdl",
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

// Media2Namespace is the WSDL namespace a device advertises in its GetServices
// response when it supports ONVIF Media2 (ver20).
const Media2Namespace = "http://www.onvif.org/ver20/media/wsdl"

// MediaCodec values as they appear in a VideoEncoder Encoding element.
type MediaCodec string

const (
	MEDIA2_CODEC_H264 MediaCodec = "H264"
	MEDIA2_CODEC_H265 MediaCodec = "H265"
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
	params          DeviceParams
	endpoints       map[string]string
	media2Supported bool
}

// Media2Supported reports whether the device advertised ONVIF Media2 (ver20)
// support in its GetServices response.
func (dev *Device) Media2Supported() bool {
	return dev.media2Supported
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

// readCapabilities decodes a GetCapabilities response. The bool reports a SOAP
// fault in place of the answer, which some cameras send under HTTP 200.
func (dev *Device) readCapabilities(resp *http.Response) (*Capabilities, bool, error) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	resp.Body.Close()

	if fault, _, err := dev.GetFault(data); err == nil && fault {
		return nil, true, nil
	}

	capabiities, err := dev.DecodeCapabilities(data)
	if err != nil {
		return nil, false, err
	}
	return capabiities, false, nil
}

// getSupportedServices registers the endpoint of each service capabilities
// gives an address for.
func (dev *Device) getSupportedServices(capabiities *Capabilities) {
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
}

// GetCapabilitiesByCategory asks for the GetCapabilities categories the driver
// uses one at a time, for a device that faults on "All": a CostarHD 4261
// (firmware 2.0.1835) answers "All", and Events alone, with "Database read
// failed" under HTTP 200, yet answers Device, Media and PTZ. A category that
// fails or faults (PTZ on a fixed camera) stays empty, so the result is never
// less than the fault gave. It registers no endpoint: NewDevice registers the
// result before the device is shared, so the driver can also call this on a
// device in use.
func (dev *Device) GetCapabilitiesByCategory(timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) *Capabilities {
	capabilities := &Capabilities{}
	for _, category := range []xsd_onvif.CapabilityCategory{"Device", "Media", "PTZ"} {
		getCapabilities := device.GetCapabilities{Category: category}
		resp, err := dev.CallMethod(getCapabilities, timeDiff, httpClient, true)
		if (err != nil || resp.StatusCode != http.StatusOK) && !omitSecurityHeader {
			if resp != nil {
				resp.Body.Close()
			}
			resp, err = dev.CallMethod(getCapabilities, timeDiff, httpClient, false)
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			if resp != nil {
				resp.Body.Close()
			}
			continue
		}

		answer, faulted, err := dev.readCapabilities(resp)
		if err != nil || faulted {
			continue
		}
		switch category {
		case "Device":
			capabilities.Device = answer.Device
		case "Media":
			capabilities.Media = answer.Media
		case "PTZ":
			capabilities.PTZ = answer.PTZ
		}
	}
	return capabilities
}

// callError is the cause of a failed CallMethod: the transport error when there
// is one, otherwise the non-OK HTTP status line. NewDevice wraps it with %w so
// callers that unwrap to the root cause see the cause alone, not the camera
// address context (and a nil err no longer renders as "%!s(<nil>)").
func callError(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp != nil {
		return fmt.Errorf("%s", resp.Status)
	}
	return fmt.Errorf("no response")
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
			return nil, time.Duration(0), nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %w", dev.params.Xaddr, callError(resp, err))
		}
	}
	timeDiff, err := dev.getTimeDiff(resp)
	if err != nil {
		return nil, time.Duration(0), nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetSystemDateAndTime): %w", dev.params.Xaddr, err)
	}

	getCapabilities := device.GetCapabilities{Category: "All"}
	resp, err = dev.CallMethod(getCapabilities, timeDiff, httpClient, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		if !omitSecurityHeader {
			resp, err = dev.CallMethod(getCapabilities, timeDiff, httpClient, false)
		}
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, timeDiff, nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services (GetCapabilities): %w", dev.params.Xaddr, callError(resp, err))
		}
	}

	capabilities, faulted, err := dev.readCapabilities(resp)
	if err != nil {
		return nil, timeDiff, nil, err
	}
	if faulted {
		capabilities = dev.GetCapabilitiesByCategory(timeDiff, httpClient, omitSecurityHeader)
	}
	dev.getSupportedServices(capabilities)

	// Probe for Media2 (ver20) support via the canonical GetServices operation.
	// This is best-effort and non-fatal: older cameras may not implement
	// GetServices at all, in which case the device simply remains Media1-only.
	dev.detectMedia2Support(timeDiff, httpClient, omitSecurityHeader)

	return dev, timeDiff, capabilities, nil
}

// detectMedia2Support calls the ONVIF GetServices operation and, if the device
// advertises the ver20/media namespace, records Media2 support and registers
// the media2 endpoint so Media2 requests can be routed to it. Any failure
// (operation unsupported, transport error, decode error) leaves the device
// Media1-only and is not propagated, since Media1 remains the fallback.
func (dev *Device) detectMedia2Support(timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) {
	getServices := device.GetServices{IncludeCapability: false}
	resp, err := dev.CallMethod(getServices, timeDiff, httpClient, omitSecurityHeader)
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil {
			resp.Body.Close()
		}
		return
	}

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	services, err := dev.DecodeServices(data)
	if err != nil {
		return
	}

	for _, service := range services {
		if service.Namespace == Media2Namespace {
			dev.media2Supported = true
			if service.XAddr != "" {
				dev.addEndpoint("media2", service.XAddr)
			}
			break
		}
	}
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
func (dev Device) CallMethod(method any, timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, err
	}
	return dev.callMethodDo(endpoint, method, timeDiff, httpClient, omitSecurityHeader)
}

// CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method any, timeDiff time.Duration, httpClient *http.Client, omitSecurityHeader bool) (*http.Response, error) {
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
