package onvif

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"slices"
	"strings"
	"sync"
	"testing"
)

const soapEnvelope = `<?xml version="1.0" encoding="UTF-8"?>` +
	`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" ` +
	`xmlns:tds="http://www.onvif.org/ver10/device/wsdl" xmlns:tt="http://www.onvif.org/ver10/schema">` +
	`<SOAP-ENV:Body>%s</SOAP-ENV:Body></SOAP-ENV:Envelope>`

// The fault a CostarHD 4261 (firmware 2.0.1835) sends under HTTP 200.
const databaseReadFault = `<SOAP-ENV:Fault>` +
	`<SOAP-ENV:Code><SOAP-ENV:Value>SOAP-ENV:Receiver</SOAP-ENV:Value></SOAP-ENV:Code>` +
	`<SOAP-ENV:Reason><SOAP-ENV:Text xml:lang="en">Database read failed</SOAP-ENV:Text></SOAP-ENV:Reason>` +
	`</SOAP-ENV:Fault>`

const systemDateAndTime = `<tds:GetSystemDateAndTimeResponse><tds:SystemDateAndTime><tt:UTCDateTime>` +
	`<tt:Time><tt:Hour>2</tt:Hour><tt:Minute>57</tt:Minute><tt:Second>2</tt:Second></tt:Time>` +
	`<tt:Date><tt:Year>2026</tt:Year><tt:Month>10</tt:Month><tt:Day>28</tt:Day></tt:Date>` +
	`</tt:UTCDateTime></tds:SystemDateAndTime></tds:GetSystemDateAndTimeResponse>`

// The camera reports its own address; addEndpoint replaces the host.
var categoryCapabilities = map[string]string{
	"Device": `<tt:Device><tt:XAddr>http://192.0.2.10/onvif/device_service</tt:XAddr></tt:Device>`,
	"Events": `<tt:Events><tt:XAddr>http://192.0.2.10/events</tt:XAddr></tt:Events>`,
	"Media":  `<tt:Media><tt:XAddr>http://192.0.2.10/onvif/media</tt:XAddr></tt:Media>`,
	"PTZ":    `<tt:PTZ><tt:XAddr>http://192.0.2.10/onvif/ptz</tt:XAddr></tt:PTZ>`,
}

var categoryPattern = regexp.MustCompile(`Category>(\w+)<`)

// fakeCamera answers what NewDevice sends, faulting (under HTTP 200) on the
// GetCapabilities categories in faults, and records every category asked.
type fakeCamera struct {
	faults map[string]bool
	mu     sync.Mutex
	asked  []string
}

func (c *fakeCamera) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request, _ := io.ReadAll(r.Body)
	body := ""
	switch {
	case strings.Contains(string(request), "GetSystemDateAndTime"):
		body = systemDateAndTime
	case strings.Contains(string(request), "GetCapabilities"):
		category := categoryPattern.FindStringSubmatch(string(request))[1]
		c.mu.Lock()
		c.asked = append(c.asked, category)
		c.mu.Unlock()
		if c.faults[category] {
			body = databaseReadFault
			break
		}
		answer := categoryCapabilities[category]
		if category == "All" {
			for _, name := range []string{"Device", "Events", "Media", "PTZ"} {
				if !c.faults[name] {
					answer += categoryCapabilities[name]
				}
			}
		}
		body = `<tds:GetCapabilitiesResponse><tds:Capabilities>` + answer + `</tds:Capabilities></tds:GetCapabilitiesResponse>`
	default:
		body = `<tds:GetServicesResponse></tds:GetServicesResponse>`
	}
	w.Header().Set("Content-Type", "application/soap+xml; charset=utf-8")
	io.WriteString(w, strings.Replace(soapEnvelope, "%s", body, 1))
}

func connect(t *testing.T, faults ...string) (*Device, *Capabilities, *fakeCamera, string) {
	t.Helper()
	camera := &fakeCamera{faults: map[string]bool{}}
	for _, category := range faults {
		camera.faults[category] = true
	}
	server := httptest.NewServer(camera)
	t.Cleanup(server.Close)
	xaddr := strings.TrimPrefix(server.URL, "http://")

	dev, _, capabilities, err := NewDevice(DeviceParams{Xaddr: xaddr}, server.Client(), true)
	if err != nil {
		t.Fatalf("NewDevice: %s", err)
	}
	return dev, capabilities, camera, xaddr
}

func endpoint(dev *Device, pkg string) string {
	url, err := dev.getEndpoint(pkg)
	if err != nil {
		return err.Error()
	}
	return url
}

// The other cameras: "All" answers, so nothing more is asked.
func TestCapabilitiesAnsweredOnAllAsksNothingMore(t *testing.T) {
	dev, capabilities, camera, xaddr := connect(t)

	if !slices.Equal(camera.asked, []string{"All"}) {
		t.Errorf("categories asked: %v, want [All]", camera.asked)
	}
	if got, want := endpoint(dev, "ptz"), "http://"+xaddr+"/onvif/ptz"; got != want {
		t.Errorf("ptz resolves to %s, want %s", got, want)
	}
	if capabilities.Events.XAddr == "" {
		t.Error("the Events capability from All was lost")
	}
}

// The reported cameras: "All" and Events fault, the categories the driver
// uses are asked one at a time and each registers its endpoint.
func TestCapabilitiesFaultOnAllAsksEachCategory(t *testing.T) {
	dev, capabilities, camera, xaddr := connect(t, "All", "Events")

	if want := []string{"All", "Device", "Media", "PTZ"}; !slices.Equal(camera.asked, want) {
		t.Errorf("categories asked: %v, want %v", camera.asked, want)
	}
	for pkg, path := range map[string]string{"device": "/onvif/device_service", "media": "/onvif/media", "ptz": "/onvif/ptz"} {
		if got, want := endpoint(dev, pkg), "http://"+xaddr+path; got != want {
			t.Errorf("%s resolves to %s, want %s", pkg, got, want)
		}
	}
	if capabilities.Media.XAddr == "" || capabilities.PTZ.XAddr == "" {
		t.Errorf("the per-category answers were not kept: %+v", capabilities)
	}
}

// The driver also calls GetCapabilitiesByCategory on a device in use, so it
// must return what it reads without registering any of it.
func TestCapabilitiesByCategoryRegistersNothing(t *testing.T) {
	dev, _, _, _ := connect(t)
	delete(dev.endpoints, "ptz")

	capabilities := dev.GetCapabilitiesByCategory(0, http.DefaultClient, true)

	if capabilities.PTZ.XAddr == "" {
		t.Error("the PTZ answer was not returned")
	}
	if got := endpoint(dev, "ptz"); got != "target endpoint service not found" {
		t.Errorf("ptz was registered: %s", got)
	}
}

// A fixed camera faulting on "All": its PTZ fault leaves PTZ unregistered
// and does not fail the connection.
func TestCapabilitiesFaultOnAllSkipsAFaultedCategory(t *testing.T) {
	dev, capabilities, _, xaddr := connect(t, "All", "PTZ")

	if got := endpoint(dev, "ptz"); got != "target endpoint service not found" {
		t.Errorf("ptz resolves to %s, want no endpoint", got)
	}
	if got, want := endpoint(dev, "media"), "http://"+xaddr+"/onvif/media"; got != want {
		t.Errorf("media resolves to %s, want %s", got, want)
	}
	if capabilities.PTZ.XAddr != "" {
		t.Errorf("PTZ capability from a fault: %+v", capabilities.PTZ)
	}
}
