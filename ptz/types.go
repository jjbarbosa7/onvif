package ptz

import (
	"github.com/jjbarbosa7/onvif/xsd"
	"github.com/jjbarbosa7/onvif/xsd_onvif"
)

type Capabilities struct {
	EFlip                       xsd.Boolean `xml:"EFlip,attr"`
	Reverse                     xsd.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus                  xsd.Boolean `xml:"MoveStatus,attr"`
	StatusPosition              xsd.Boolean `xml:"StatusPosition,attr"`
}

//PTZ main types

type GetServiceCapabilities struct {
	XMLName string `xml:"tptz:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetNodes struct {
	XMLName string `xml:"tptz:GetNodes"`
}

type GetNodesResponse struct {
	PTZNode xsd_onvif.PTZNode
}

type GetNode struct {
	XMLName   string                   `xml:"tptz:GetNode"`
	NodeToken xsd_onvif.ReferenceToken `xml:"tptz:NodeToken"`
}

type GetNodeResponse struct {
	PTZNode xsd_onvif.PTZNode
}

type GetConfiguration struct {
	XMLName      string                   `xml:"tptz:GetConfiguration"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetConfigurationResponse struct {
	PTZConfiguration xsd_onvif.PTZConfiguration
}

type GetConfigurations struct {
	XMLName string `xml:"tptz:GetConfigurations"`
}

type GetConfigurationsResponse struct {
	PTZConfiguration xsd_onvif.PTZConfiguration
}

type SetConfiguration struct {
	XMLName          string                     `xml:"tptz:SetConfiguration"`
	PTZConfiguration xsd_onvif.PTZConfiguration `xml:"tptz:PTZConfiguration"`
	ForcePersistence xsd.Boolean                `xml:"tptz:ForcePersistence"`
}

type SetConfigurationResponse struct {
}

type GetConfigurationOptions struct {
	XMLName      string                   `xml:"tptz:GetConfigurationOptions"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions xsd_onvif.PTZConfigurationOptions
}

type SendAuxiliaryCommand struct {
	XMLName       string                   `xml:"tptz:SendAuxiliaryCommand"`
	ProfileToken  xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	AuxiliaryData xsd_onvif.AuxiliaryData  `xml:"tptz:AuxiliaryData"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse xsd_onvif.AuxiliaryData
}

type GetPresets struct {
	XMLName      string                   `xml:"tptz:GetPresets"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetPresetsResponse struct {
	Preset []xsd_onvif.PTZPreset
}

type SetPreset struct {
	XMLName      string                   `xml:"tptz:SetPreset"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetName   xsd.String               `xml:"tptz:PresetName"`
	PresetToken  xsd_onvif.ReferenceToken `xml:"tptz:PresetToken,omitempty"`
}

type SetPresetResponse struct {
	PresetToken xsd_onvif.ReferenceToken
}

type RemovePreset struct {
	XMLName      string                   `xml:"tptz:RemovePreset"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetToken  xsd_onvif.ReferenceToken `xml:"tptz:PresetToken"`
}

type RemovePresetResponse struct {
}

type GotoPreset struct {
	XMLName      string                   `xml:"tptz:GotoPreset"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetToken  xsd_onvif.ReferenceToken `xml:"tptz:PresetToken"`
	Speed        xsd_onvif.PTZSpeed       `xml:"tptz:Speed"`
}

type GotoPresetResponse struct {
}

type GotoHomePosition struct {
	XMLName      string                   `xml:"tptz:GotoHomePosition"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Speed        xsd_onvif.PTZSpeed       `xml:"tptz:Speed"`
}

type GotoHomePositionResponse struct {
}

type SetHomePosition struct {
	XMLName      string                   `xml:"tptz:SetHomePosition"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type SetHomePositionResponse struct {
}

type ContinuousMove struct {
	XMLName      string                   `xml:"tptz:ContinuousMove"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Velocity     xsd_onvif.PTZSpeed       `xml:"tptz:Velocity"`
	Timeout      xsd.Duration             `xml:"tptz:Timeout"`
}

type ContinuousMoveResponse struct {
}

type RelativeMove struct {
	XMLName      string                   `xml:"tptz:RelativeMove"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Translation  xsd_onvif.PTZVector      `xml:"tptz:Translation"`
	Speed        xsd_onvif.PTZSpeed       `xml:"tptz:Speed"`
}

type RelativeMoveResponse struct {
}

type GetStatus struct {
	XMLName      string                   `xml:"tptz:GetStatus"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetStatusResponse struct {
	PTZStatus xsd_onvif.PTZStatus
}

type AbsoluteMove struct {
	XMLName      string                   `xml:"tptz:AbsoluteMove"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Position     xsd_onvif.PTZVector      `xml:"tptz:Position"`
	Speed        xsd_onvif.PTZSpeed       `xml:"tptz:Speed"`
}

type AbsoluteMoveResponse struct {
}

type GeoMove struct {
	XMLName      string                   `xml:"tptz:GeoMove"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Target       xsd_onvif.GeoLocation    `xml:"tptz:Target"`
	Speed        xsd_onvif.PTZSpeed       `xml:"tptz:Speed"`
	AreaHeight   xsd.Float                `xml:"tptz:AreaHeight"`
	AreaWidth    xsd.Float                `xml:"tptz:AreaWidth"`
}

type GeoMoveResponse struct {
}

type Stop struct {
	XMLName      string                   `xml:"tptz:Stop"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PanTilt      xsd.Boolean              `xml:"tptz:PanTilt"`
	Zoom         xsd.Boolean              `xml:"tptz:Zoom"`
}

type StopResponse struct {
}

type GetPresetTours struct {
	XMLName      string                   `xml:"tptz:GetPresetTours"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetPresetToursResponse struct {
	PresetTour xsd_onvif.PresetTour
}

type GetPresetTour struct {
	XMLName         string                   `xml:"tptz:GetPresetTour"`
	ProfileToken    xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken xsd_onvif.ReferenceToken `xml:"tptz:PresetTourToken"`
}

type GetPresetTourResponse struct {
	PresetTour xsd_onvif.PresetTour
}

type GetPresetTourOptions struct {
	XMLName         string                   `xml:"tptz:GetPresetTourOptions"`
	ProfileToken    xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken xsd_onvif.ReferenceToken `xml:"tptz:PresetTourToken"`
}

type GetPresetTourOptionsResponse struct {
	Options xsd_onvif.PTZPresetTourOptions
}

type CreatePresetTour struct {
	XMLName      string                   `xml:"tptz:CreatePresetTour"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type CreatePresetTourResponse struct {
	PresetTourToken xsd_onvif.ReferenceToken
}

type ModifyPresetTour struct {
	XMLName      string                   `xml:"tptz:ModifyPresetTour"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTour   xsd_onvif.PresetTour     `xml:"tptz:PresetTour"`
}

type ModifyPresetTourResponse struct {
}

type OperatePresetTour struct {
	XMLName         string                           `xml:"tptz:OperatePresetTour"`
	ProfileToken    xsd_onvif.ReferenceToken         `xml:"tptz:ProfileToken"`
	PresetTourToken xsd_onvif.ReferenceToken         `xml:"onvif:PresetTourToken"`
	Operation       xsd_onvif.PTZPresetTourOperation `xml:"onvif:Operation"`
}

type OperatePresetTourResponse struct {
}

type RemovePresetTour struct {
	XMLName         string                   `xml:"tptz:RemovePresetTour"`
	ProfileToken    xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken xsd_onvif.ReferenceToken `xml:"tptz:PresetTourToken"`
}

type RemovePresetTourResponse struct {
}

type GetCompatibleConfigurations struct {
	XMLName      string                   `xml:"tptz:GetCompatibleConfigurations"`
	ProfileToken xsd_onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}

type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration xsd_onvif.PTZConfiguration
}
