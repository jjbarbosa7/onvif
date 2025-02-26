package imaging

import (
	"github.com/jjbarbosa7/onvif/xsd"
	"github.com/jjbarbosa7/onvif/xsd_onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string                   `xml:"timg:GetImagingSettings"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                      `xml:"timg:SetImagingSettings"`
	VideoSourceToken xsd_onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	ImagingSettings  xsd_onvif.ImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                 `xml:"timg:ForcePersistence"`
}

type GetOptions struct {
	XMLName          string                   `xml:"timg:GetOptions"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Move struct {
	XMLName          string                   `xml:"timg:Move"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            xsd_onvif.FocusMove      `xml:"timg:Focus"`
}

type GetMoveOptions struct {
	XMLName          string                   `xml:"timg:GetMoveOptions"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Stop struct {
	XMLName          string                   `xml:"timg:Stop"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string                   `xml:"timg:GetStatus"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string                   `xml:"timg:GetPresets"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string                   `xml:"timg:GetCurrentPreset"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string                   `xml:"timg:SetCurrentPreset"`
	VideoSourceToken xsd_onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      xsd_onvif.ReferenceToken `xml:"timg:PresetToken"`
}
