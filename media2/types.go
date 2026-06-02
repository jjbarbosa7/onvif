// Package media2 contains request types for the ONVIF Media2 (ver20) service
// (namespace http://www.onvif.org/ver20/media/wsdl, prefix tr2).
//
// Request types must live in this package because Device.CallMethod routes a
// request to a service endpoint by reflecting the request struct's package name
// (lowercased) to an endpoint key. NewDevice registers the Media2 endpoint under
// the key "media2", so these types resolve there. Responses are decoded via
// etree (see onvif.DecodeProfiles2 et al.), so no response structs are defined.
package media2

import (
	"github.com/jjbarbosa7/onvif/xsd_onvif"
)

// GetProfiles retrieves the media profiles available on the camera via Media2,
// including H265 profiles, which the Media1 GetProfiles response omits entirely.
//
// Unlike Media1, Media2 does NOT embed configurations by default: the returned
// profiles carry an empty <Configurations> unless one or more Type values are
// requested (tr2:ConfigurationEnumeration). "All" requests every configuration
// type, so the VideoEncoder (incl. its encoding/resolution/etc.) is populated.
// Token is optional; omitting it returns every profile.
type GetProfiles struct {
	XMLName string                   `xml:"tr2:GetProfiles"`
	Token   xsd_onvif.ReferenceToken `xml:"tr2:Token,omitempty"`
	Type    string                   `xml:"tr2:Type,omitempty"`
}
