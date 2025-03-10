package device

import (
	"github.com/jjbarbosa7/onvif/xsd"
	"github.com/jjbarbosa7/onvif/xsd_onvif"
)

type SetSystemDateAndTime struct {
	XMLName         string                    `xml:"tds:SetSystemDateAndTime"`
	DateTimeType    xsd_onvif.SetDateTimeType `xml:"tds:DateTimeType"`
	DaylightSavings xsd.Boolean               `xml:"tds:DaylightSavings"`
	TimeZone        xsd_onvif.TimeZone        `xml:"tds:TimeZone"`
	UTCDateTime     xsd_onvif.DateTime        `xml:"tds:UTCDateTime"`
}

type SetSystemDateAndTimeResponse struct {
}

type GetSystemDateAndTime struct {
	XMLName string `xml:"tds:GetSystemDateAndTime"`
}

type GetDeviceInformation struct {
	XMLName string `xml:"tds:GetDeviceInformation"`
}

type GetDeviceInformationResponse struct {
	DevideInformation xsd_onvif.DeviceInformation
}

type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime xsd_onvif.SystemDateTime
}

type SetSystemFactoryDefault struct {
	XMLName        string                       `xml:"tds:SetSystemFactoryDefault"`
	FactoryDefault xsd_onvif.FactoryDefaultType `xml:"tds:FactoryDefault"`
}

type SetSystemFactoryDefaultResponse struct {
}

type UpgradeSystemFirmware struct {
	XMLName  string                   `xml:"tds:UpgradeSystemFirmware"`
	Firmware xsd_onvif.AttachmentData `xml:"tds:Firmware"`
}

type UpgradeSystemFirmwareResponse struct {
	Message string
}

type SystemReboot struct {
	XMLName string `xml:"tds:SystemReboot"`
}

type SystemRebootResponse struct {
	Message string
}

// TODO: one or more repetitions
type RestoreSystem struct {
	XMLName     string               `xml:"tds:RestoreSystem"`
	BackupFiles xsd_onvif.BackupFile `xml:"tds:BackupFiles"`
}

type RestoreSystemResponse struct {
}

type GetSystemBackup struct {
	XMLName string `xml:"tds:GetSystemBackup"`
}

type GetSystemBackupResponse struct {
	BackupFiles xsd_onvif.BackupFile
}

type GetSystemLog struct {
	XMLName string                  `xml:"tds:GetSystemLog"`
	LogType xsd_onvif.SystemLogType `xml:"tds:LogType"`
}

type GetSystemLogResponse struct {
	SystemLog xsd_onvif.SystemLog
}

type GetSystemSupportInformation struct {
	XMLName string `xml:"tds:GetSystemSupportInformation"`
}

type GetSystemSupportInformationResponse struct {
	SupportInformation xsd_onvif.SupportInformation
}

type GetScopes struct {
	XMLName string `xml:"tds:GetScopes"`
}

type GetScopesResponse struct {
	Scopes xsd_onvif.Scope
}

// TODO: one or more scopes
type SetScopes struct {
	XMLName string     `xml:"tds:SetScopes"`
	Scopes  xsd.AnyURI `xml:"tds:Scopes"`
}

type SetScopesResponse struct {
}

// TODO: list of scopes
type AddScopes struct {
	XMLName   string     `xml:"tds:AddScopes"`
	ScopeItem xsd.AnyURI `xml:"tds:ScopeItem"`
}

type AddScopesResponse struct {
}

// TODO: One or more repetitions
type RemoveScopes struct {
	XMLName   string     `xml:"tds:RemoveScopes"`
	ScopeItem xsd.AnyURI `xml:"onvif:ScopeItem"`
}

type RemoveScopesResponse struct {
	ScopeItem xsd.AnyURI
}

type GetDiscoveryMode struct {
	XMLName string `xml:"tds:GetDiscoveryMode"`
}

type GetDiscoveryModeResponse struct {
	DiscoveryMode xsd_onvif.DiscoveryMode
}

type SetDiscoveryMode struct {
	XMLName       string                  `xml:"tds:SetDiscoveryMode"`
	DiscoveryMode xsd_onvif.DiscoveryMode `xml:"tds:DiscoveryMode"`
}

type SetDiscoveryModeResponse struct {
}

type GetRemoteDiscoveryMode struct {
	XMLName string `xml:"tds:GetRemoteDiscoveryMode"`
}

type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode xsd_onvif.DiscoveryMode
}

type SetRemoteDiscoveryMode struct {
	XMLName             string                  `xml:"tds:SetRemoteDiscoveryMode"`
	RemoteDiscoveryMode xsd_onvif.DiscoveryMode `xml:"tds:RemoteDiscoveryMode"`
}

type SetRemoteDiscoveryModeResponse struct {
}

type GetDPAddresses struct {
	XMLName string `xml:"tds:GetDPAddresses"`
}

type GetDPAddressesResponse struct {
	DPAddress xsd_onvif.NetworkHost
}

type SetDPAddresses struct {
	XMLName   string                `xml:"tds:SetDPAddresses"`
	DPAddress xsd_onvif.NetworkHost `xml:"tds:DPAddress"`
}

type SetDPAddressesResponse struct {
}

type GetEndpointReference struct {
	XMLName string `xml:"tds:GetEndpointReference"`
}

type GetEndpointReferenceResponse struct {
	GUID string
}

type GetRemoteUser struct {
	XMLName string `xml:"tds:GetRemoteUser"`
}

type GetRemoteUserResponse struct {
	RemoteUser xsd_onvif.RemoteUser
}

type SetRemoteUser struct {
	XMLName    string               `xml:"tds:SetRemoteUser"`
	RemoteUser xsd_onvif.RemoteUser `xml:"tds:RemoteUser"`
}

type SetRemoteUserResponse struct {
}

type GetUsers struct {
	XMLName string `xml:"tds:GetUsers"`
}

type GetUsersResponse struct {
	User xsd_onvif.User
}

// TODO: List of users
type CreateUsers struct {
	XMLName string         `xml:"tds:CreateUsers"`
	User    xsd_onvif.User `xml:"tds:User,omitempty"`
}

type CreateUsersResponse struct {
}

// TODO: one or more Username
type DeleteUsers struct {
	XMLName  xsd.String `xml:"tds:DeleteUsers"`
	Username xsd.String `xml:"tds:Username"`
}

type DeleteUsersResponse struct {
}

type SetUser struct {
	XMLName string         `xml:"tds:SetUser"`
	User    xsd_onvif.User `xml:"tds:User"`
}

type SetUserResponse struct {
}

type GetWsdlUrl struct {
	XMLName string `xml:"tds:GetWsdlUrl"`
}

type GetWsdlUrlResponse struct {
	WsdlUrl xsd.AnyURI
}

type GetCapabilities struct {
	XMLName  string                       `xml:"tds:GetCapabilities"`
	Category xsd_onvif.CapabilityCategory `xml:"tds:Category"`
}

type GetCapabilitiesResponse struct {
	Capabilities xsd_onvif.Capabilities
}

type GetHostname struct {
	XMLName string `xml:"tds:GetHostname"`
}

type GetHostnameResponse struct {
	HostnameInformation xsd_onvif.HostnameInformation
}

type SetHostname struct {
	XMLName string    `xml:"tds:SetHostname"`
	Name    xsd.Token `xml:"tds:Name"`
}

type SetHostnameResponse struct {
}

type SetHostnameFromDHCP struct {
	XMLName  string      `xml:"tds:SetHostnameFromDHCP"`
	FromDHCP xsd.Boolean `xml:"tds:FromDHCP"`
}

type SetHostnameFromDHCPResponse struct {
	RebootNeeded xsd.Boolean
}

type GetDNS struct {
	XMLName string `xml:"tds:GetDNS"`
}

type GetDNSResponse struct {
	DNSInformation xsd_onvif.DNSInformation
}

type SetDNS struct {
	XMLName      string              `xml:"tds:SetDNS"`
	FromDHCP     xsd.Boolean         `xml:"tds:FromDHCP"`
	SearchDomain xsd.Token           `xml:"tds:SearchDomain"`
	DNSManual    xsd_onvif.IPAddress `xml:"tds:DNSManual"`
}

type SetDNSResponse struct {
}

type GetNTP struct {
	XMLName string `xml:"tds:GetNTP"`
}

type GetNTPResponse struct {
	NTPInformation xsd_onvif.NTPInformation
}

type SetNTP struct {
	XMLName   string                `xml:"tds:SetNTP"`
	FromDHCP  xsd.Boolean           `xml:"tds:FromDHCP"`
	NTPManual xsd_onvif.NetworkHost `xml:"tds:NTPManual"`
}

type SetNTPResponse struct {
}

type GetDynamicDNS struct {
	XMLName string `xml:"tds:GetDynamicDNS"`
}

type GetDynamicDNSResponse struct {
	DynamicDNSInformation xsd_onvif.DynamicDNSInformation
}

type SetDynamicDNS struct {
	XMLName string                   `xml:"tds:SetDynamicDNS"`
	Type    xsd_onvif.DynamicDNSType `xml:"tds:Type"`
	Name    xsd_onvif.DNSName        `xml:"tds:Name"`
	TTL     xsd.Duration             `xml:"tds:TTL"`
}

type SetDynamicDNSResponse struct {
}

type GetNetworkInterfaces struct {
	XMLName string `xml:"tds:GetNetworkInterfaces"`
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces xsd_onvif.NetworkInterface
}

type SetNetworkInterfaces struct {
	XMLName          string                                     `xml:"tds:SetNetworkInterfaces"`
	InterfaceToken   xsd_onvif.ReferenceToken                   `xml:"tds:InterfaceToken"`
	NetworkInterface xsd_onvif.NetworkInterfaceSetConfiguration `xml:"tds:NetworkInterface"`
}

type SetNetworkInterfacesResponse struct {
	RebootNeeded xsd.Boolean
}

type GetNetworkProtocols struct {
	XMLName string `xml:"tds:GetNetworkProtocols"`
}

type GetNetworkProtocolsResponse struct {
	NetworkProtocols xsd_onvif.NetworkProtocol
}

type SetNetworkProtocols struct {
	XMLName          string                    `xml:"tds:SetNetworkProtocols"`
	NetworkProtocols xsd_onvif.NetworkProtocol `xml:"tds:NetworkProtocols"`
}

type SetNetworkProtocolsResponse struct {
}

type GetNetworkDefaultGateway struct {
	XMLName string `xml:"tds:GetNetworkDefaultGateway"`
}

type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway xsd_onvif.NetworkGateway
}

type SetNetworkDefaultGateway struct {
	XMLName     string                `xml:"tds:SetNetworkDefaultGateway"`
	IPv4Address xsd_onvif.IPv4Address `xml:"tds:IPv4Address"`
	IPv6Address xsd_onvif.IPv6Address `xml:"tds:IPv6Address"`
}

type SetNetworkDefaultGatewayResponse struct {
}

type GetZeroConfiguration struct {
	XMLName string `xml:"tds:GetZeroConfiguration"`
}

type GetZeroConfigurationResponse struct {
	ZeroConfiguration xsd_onvif.NetworkZeroConfiguration
}

type SetZeroConfiguration struct {
	XMLName        string                   `xml:"tds:SetZeroConfiguration"`
	InterfaceToken xsd_onvif.ReferenceToken `xml:"tds:InterfaceToken"`
	Enabled        xsd.Boolean              `xml:"tds:Enabled"`
}

type SetZeroConfigurationResponse struct {
}

type GetIPAddressFilter struct {
	XMLName string `xml:"tds:GetIPAddressFilter"`
}

type GetIPAddressFilterResponse struct {
	IPAddressFilter xsd_onvif.IPAddressFilter
}

type SetIPAddressFilter struct {
	XMLName         string                    `xml:"tds:SetIPAddressFilter"`
	IPAddressFilter xsd_onvif.IPAddressFilter `xml:"tds:IPAddressFilter"`
}

type SetIPAddressFilterResponse struct {
}

// This operation adds an IP filter address to a device.
// If the device supports device access control based on
// IP filtering rules (denied or accepted ranges of IP addresses),
// the device shall support adding of IP filtering addresses through
// the AddIPAddressFilter command.
type AddIPAddressFilter struct {
	XMLName         string                    `xml:"tds:AddIPAddressFilter"`
	IPAddressFilter xsd_onvif.IPAddressFilter `xml:"tds:IPAddressFilter"`
}

type AddIPAddressFilterResponse struct {
}

type RemoveIPAddressFilter struct {
	XMLName         string                    `xml:"tds:RemoveIPAddressFilter"`
	IPAddressFilter xsd_onvif.IPAddressFilter `xml:"onvif:IPAddressFilter"`
}

type RemoveIPAddressFilterResponse struct {
}

type GetAccessPolicy struct {
	XMLName string `xml:"tds:GetAccessPolicy"`
}

type GetAccessPolicyResponse struct {
	PolicyFile xsd_onvif.BinaryData
}

type SetAccessPolicy struct {
	XMLName    string               `xml:"tds:SetAccessPolicy"`
	PolicyFile xsd_onvif.BinaryData `xml:"tds:PolicyFile"`
}

type SetAccessPolicyResponse struct {
}

type CreateCertificate struct {
	XMLName        string       `xml:"tds:CreateCertificate"`
	CertificateID  xsd.Token    `xml:"tds:CertificateID,omitempty"`
	Subject        string       `xml:"tds:Subject,omitempty"`
	ValidNotBefore xsd.DateTime `xml:"tds:ValidNotBefore,omitempty"`
	ValidNotAfter  xsd.DateTime `xml:"tds:ValidNotAfter,omitempty"`
}

type CreateCertificateResponse struct {
	NvtCertificate xsd_onvif.Certificate
}

type GetCertificates struct {
	XMLName string `xml:"tds:GetCertificates"`
}

type GetCertificatesResponse struct {
	NvtCertificate xsd_onvif.Certificate
}

type GetCertificatesStatus struct {
	XMLName string `xml:"tds:GetCertificatesStatus"`
}

type GetCertificatesStatusResponse struct {
	CertificateStatus xsd_onvif.CertificateStatus
}

type SetCertificatesStatus struct {
	XMLName           string                      `xml:"tds:SetCertificatesStatus"`
	CertificateStatus xsd_onvif.CertificateStatus `xml:"tds:CertificateStatus"`
}

type SetCertificatesStatusResponse struct {
}

// TODO: List of CertificateID
type DeleteCertificates struct {
	XMLName       string    `xml:"tds:DeleteCertificates"`
	CertificateID xsd.Token `xml:"tds:CertificateID"`
}

type DeleteCertificatesResponse struct {
}

// TODO: Откуда onvif:data = cid:21312413412
type GetPkcs10Request struct {
	XMLName       string               `xml:"tds:GetPkcs10Request"`
	CertificateID xsd.Token            `xml:"tds:CertificateID"`
	Subject       xsd.String           `xml:"tds:Subject"`
	Attributes    xsd_onvif.BinaryData `xml:"tds:Attributes"`
}

type GetPkcs10RequestResponse struct {
	Pkcs10Request xsd_onvif.BinaryData
}

// TODO: one or more NTVCertificate
type LoadCertificates struct {
	XMLName        string                `xml:"tds:LoadCertificates"`
	NVTCertificate xsd_onvif.Certificate `xml:"tds:NVTCertificate"`
}

type LoadCertificatesResponse struct {
}

type GetClientCertificateMode struct {
	XMLName string `xml:"tds:GetClientCertificateMode"`
}

type GetClientCertificateModeResponse struct {
	Enabled xsd.Boolean
}

type SetClientCertificateMode struct {
	XMLName string      `xml:"tds:SetClientCertificateMode"`
	Enabled xsd.Boolean `xml:"tds:Enabled"`
}

type SetClientCertificateModeResponse struct {
}

type GetRelayOutputs struct {
	XMLName string `xml:"tds:GetRelayOutputs"`
}

type GetRelayOutputsResponse struct {
	RelayOutputs xsd_onvif.RelayOutput
}

type SetRelayOutputSettings struct {
	XMLName          string                        `xml:"tds:SetRelayOutputSettings"`
	RelayOutputToken xsd_onvif.ReferenceToken      `xml:"tds:RelayOutputToken"`
	Properties       xsd_onvif.RelayOutputSettings `xml:"tds:Properties"`
}

type SetRelayOutputSettingsResponse struct {
}

type SetRelayOutputState struct {
	XMLName          string                      `xml:"tds:SetRelayOutputState"`
	RelayOutputToken xsd_onvif.ReferenceToken    `xml:"tds:RelayOutputToken"`
	LogicalState     xsd_onvif.RelayLogicalState `xml:"tds:LogicalState"`
}

type SetRelayOutputStateResponse struct {
}

type SendAuxiliaryCommand struct {
	XMLName          string                  `xml:"tds:SendAuxiliaryCommand"`
	AuxiliaryCommand xsd_onvif.AuxiliaryData `xml:"tds:AuxiliaryCommand"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse xsd_onvif.AuxiliaryData
}

type GetCACertificates struct {
	XMLName string `xml:"tds:GetCACertificates"`
}

type GetCACertificatesResponse struct {
	CACertificate xsd_onvif.Certificate
}

// TODO: one or more CertificateWithPrivateKey
type LoadCertificateWithPrivateKey struct {
	XMLName                   string                              `xml:"tds:LoadCertificateWithPrivateKey"`
	CertificateWithPrivateKey xsd_onvif.CertificateWithPrivateKey `xml:"tds:CertificateWithPrivateKey"`
}

type LoadCertificateWithPrivateKeyResponse struct {
}

type GetCertificateInformation struct {
	XMLName       string    `xml:"tds:GetCertificateInformation"`
	CertificateID xsd.Token `xml:"tds:CertificateID"`
}

type GetCertificateInformationResponse struct {
	CertificateInformation xsd_onvif.CertificateInformation
}

type LoadCACertificates struct {
	XMLName       string                `xml:"tds:LoadCACertificates"`
	CACertificate xsd_onvif.Certificate `xml:"tds:CACertificate"`
}

type LoadCACertificatesResponse struct {
}

type CreateDot1XConfiguration struct {
	XMLName            string                       `xml:"tds:CreateDot1XConfiguration"`
	Dot1XConfiguration xsd_onvif.Dot1XConfiguration `xml:"tds:Dot1XConfiguration"`
}

type CreateDot1XConfigurationResponse struct {
}

type SetDot1XConfiguration struct {
	XMLName            string                       `xml:"tds:SetDot1XConfiguration"`
	Dot1XConfiguration xsd_onvif.Dot1XConfiguration `xml:"tds:Dot1XConfiguration"`
}

type SetDot1XConfigurationResponse struct {
}

type GetDot1XConfiguration struct {
	XMLName                 string                   `xml:"tds:GetDot1XConfiguration"`
	Dot1XConfigurationToken xsd_onvif.ReferenceToken `xml:"tds:Dot1XConfigurationToken"`
}

type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration xsd_onvif.Dot1XConfiguration
}

type GetDot1XConfigurations struct {
	XMLName string `xml:"tds:GetDot1XConfigurations"`
}

type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration xsd_onvif.Dot1XConfiguration
}

// TODO: Zero or more Dot1XConfigurationToken
type DeleteDot1XConfiguration struct {
	XMLName                 string                   `xml:"tds:DeleteDot1XConfiguration"`
	Dot1XConfigurationToken xsd_onvif.ReferenceToken `xml:"tds:Dot1XConfigurationToken"`
}

type DeleteDot1XConfigurationResponse struct {
}

type GetDot11Capabilities struct {
	XMLName string `xml:"tds:GetDot11Capabilities"`
}

type GetDot11CapabilitiesResponse struct {
	Capabilities xsd_onvif.Dot11Capabilities
}

type GetDot11Status struct {
	XMLName        string                   `xml:"tds:GetDot11Status"`
	InterfaceToken xsd_onvif.ReferenceToken `xml:"tds:InterfaceToken"`
}

type GetDot11StatusResponse struct {
	Status xsd_onvif.Dot11Status
}

type ScanAvailableDot11Networks struct {
	XMLName        string                   `xml:"tds:ScanAvailableDot11Networks"`
	InterfaceToken xsd_onvif.ReferenceToken `xml:"tds:InterfaceToken"`
}

type ScanAvailableDot11NetworksResponse struct {
	Networks xsd_onvif.Dot11AvailableNetworks
}

type GetSystemUris struct {
	XMLName string `xml:"tds:GetSystemUris"`
}

type GetSystemUrisResponse struct {
	SystemLogUris   xsd_onvif.SystemLogUriList
	SupportInfoUri  xsd.AnyURI
	SystemBackupUri xsd.AnyURI
	Extension       xsd.AnyType
}

type StartFirmwareUpgrade struct {
	XMLName string `xml:"tds:StartFirmwareUpgrade"`
}

type StartFirmwareUpgradeResponse struct {
	UploadUri        xsd.AnyURI
	UploadDelay      xsd.Duration
	ExpectedDownTime xsd.Duration
}

type StartSystemRestore struct {
	XMLName string `xml:"tds:StartSystemRestore"`
}

type StartSystemRestoreResponse struct {
	UploadUri        xsd.AnyURI
	ExpectedDownTime xsd.Duration
}

type GetGeoLocation struct {
	XMLName string `xml:"tds:GetGeoLocation"`
}

type GetGeoLocationResponse struct {
	Location xsd_onvif.LocationEntity
}

// TODO: one or more Location
type SetGeoLocation struct {
	XMLName  string                   `xml:"tds:SetGeoLocation"`
	Location xsd_onvif.LocationEntity `xml:"tds:Location"`
}

type SetGeoLocationResponse struct {
}

type DeleteGeoLocation struct {
	XMLName  string                   `xml:"tds:DeleteGeoLocation"`
	Location xsd_onvif.LocationEntity `xml:"tds:Location"`
}

type DeleteGeoLocationResponse struct {
}
