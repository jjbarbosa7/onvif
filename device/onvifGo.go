package device

import (
	"time"

	"github.com/jjbarbosa7/onvif/xsd/onvif"
)

// Translation of onvif types from wsdl to go

type DeviceEntity struct {
	Token onvif.ReferenceToken
}

type OnvifVersion struct {
	Major int
	Minor int
}

type SystemDateTime struct {
	DateTimeType    string
	DaylightSavings bool
	TimeZone        string
	UTCDateTime     time.Time
	LocalDateTime   time.Time
	Extension       map[string]string
}

type IntRectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

type IntRectangleRange struct {
	XRange      IntRange
	YRange      IntRange
	WidthRange  IntRange
	HeightRange IntRange
}

type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64
	Max float64
}

type OSDConfiguration struct {
	DeviceEntity
	VideoSourceConfigurationToken onvif.ReferenceToken
	Type                          string
	Position                      OSDPosConfiguration
	TextString                    OSDTextConfiguration
	Image                         OSDImgConfiguration
	Extension                     map[string]string
}

type OSDPosConfiguration struct {
	Type      string
	Pos       Vector
	Extension map[string]string
}

type Vector struct {
	X float64
	Y float64
}

type OSDTextConfiguration struct {
	IsPersistentText bool

	Type            string
	DateFormat      string
	TimeFormat      string
	FontSize        int
	FontColor       OSDColor
	BackgroundColor OSDColor
	PlainText       string
	Extension       map[string]string
}

type OSDColor struct {
	Transparent int

	Color Color
}

type Color struct {
	X          float64
	Y          float64
	Z          float64
	Colorspace string
}

type OSDImgConfiguration struct {
	ImgPath   string
	Extension map[string]string
}

type VideoSource struct {
	DeviceEntity
	Framerate  float64
	Resolution VideoResolution
	Imaging    ImagingSettings
	Extension  VideoSourceExtension
}

type VideoResolution struct {
	Width  int
	Height int
}

type ImagingSettings struct {
	BacklightCompensation BacklightCompensation
	Brightness            float64
	ColorSaturation       float64
	Contrast              float64
	Exposure              Exposure
	Focus                 FocusConfiguration
	IrCutFilter           string
	Sharpness             float64
	WideDynamicRange      WideDynamicRange
	WhiteBalance          WhiteBalance
	Extension             map[string]string
}

type BacklightCompensation struct {
	Mode  string
	Level float64
}

type Exposure struct {
	Mode            string
	Priority        string
	Window          Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain         float64
	MaxGain         float64
	MinIris         float64
	MaxIris         float64
	ExposureTime    float64
	Gain            float64
	Iris            float64
}

type Rectangle struct {
	Bottom float64
	Top    float64
	Right  float64
	Left   float64
}

type FocusConfiguration struct {
	AutoFocusMode string
	DefaultSpeed  float64
	NearLimit     float64
	FarLimit      float64
}

type WideDynamicRange struct {
	Mode  string
	Level float64
}

type WhiteBalance struct {
	Mode   string
	CrGain float64
	CbGain float64
}

type VideoSourceExtension struct {
	Imaging   ImagingSettings20
	Extension map[string]string
}

type ImagingSettings20 struct {
	BacklightCompensation *BacklightCompensation20
	Brightness            float64
	ColorSaturation       float64
	Contrast              float64
	Exposure              *Exposure20
	Focus                 *FocusConfiguration20
	IrCutFilter           *string
	Sharpness             float64
	WideDynamicRange      *WideDynamicRange20
	WhiteBalance          *WhiteBalance20
	Extension             *ImagingSettingsExtension20
}

type BacklightCompensation20 struct {
	Mode  string
	Level float64
}

type Exposure20 struct {
	Mode            string
	Priority        string
	Window          Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain         float64
	MaxGain         float64
	MinIris         float64
	MaxIris         float64
	ExposureTime    float64
	Gain            float64
	Iris            float64
}

type FocusConfiguration20 struct {
	AutoFocusMode string
	DefaultSpeed  float64
	NearLimit     float64
	FarLimit      float64
	Extension     map[string]string
}

type WideDynamicRange20 struct {
	Mode  string
	Level float64
}

type WhiteBalance20 struct {
	Mode      string
	CrGain    float64
	CbGain    float64
	Extension map[string]string
}

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization
	Extension          ImagingSettingsExtension202
}

type ImageStabilization struct {
	Mode      string
	Level     float64
	Extension map[string]string
}

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment
	Extension                 ImagingSettingsExtension203
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string
	BoundaryOffset float64
	ResponseTime   time.Duration
	Extension      map[string]string
}

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation
	Defogging        Defogging
	NoiseReduction   NoiseReduction
	Extension        map[string]string
}

type ToneCompensation struct {
	Mode      string
	Level     float64
	Extension map[string]string
}

type Defogging struct {
	Mode      string
	Level     float64
	Extension map[string]string
}

type NoiseReduction struct {
	Level float64
}

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token                       onvif.ReferenceToken
	Fixed                       bool
	Name                        string
	VideoSourceConfiguration    VideoSourceConfiguration
	AudioSourceConfiguration    AudioSourceConfiguration
	VideoEncoderConfiguration   VideoEncoderConfiguration
	AudioEncoderConfiguration   AudioEncoderConfiguration
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration
	PTZConfiguration            PTZConfiguration
	MetadataConfiguration       MetadataConfiguration
	Extension                   ProfileExtension
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode    string
	SourceToken onvif.ReferenceToken
	Bounds      IntRectangle
	Extension   VideoSourceConfigurationExtension
}

type ConfigurationEntity struct {
	Token    onvif.ReferenceToken
	Name     string
	UseCount int
}

type VideoSourceConfigurationExtension struct {
	Rotate    Rotate
	Extension VideoSourceConfigurationExtension2
}

type Rotate struct {
	Mode      string
	Degree    int
	Extension map[string]string
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription  LensDescription
	SceneOrientation SceneOrientation
}

type LensDescription struct {
	FocalLength float64
	Offset      LensOffset
	Projection  LensProjection
	XFactor     float64
}

type LensOffset struct {
	X float64
	Y float64
}

type LensProjection struct {
	Angle         float64
	Radius        float64
	Transmittance float64
}

type SceneOrientation struct {
	Mode        string
	Orientation string
}

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken onvif.ReferenceToken
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       string
	Resolution     VideoResolution
	Quality        float64
	RateControl    VideoRateControl
	MPEG4          Mpeg4Configuration
	H264           H264Configuration
	Multicast      MulticastConfiguration
	SessionTimeout time.Duration
}

type VideoRateControl struct {
	FrameRateLimit   int
	EncodingInterval int
	BitrateLimit     int
}

type Mpeg4Configuration struct {
	GovLength    int
	Mpeg4Profile string
}

type H264Configuration struct {
	GovLength   int
	H264Profile string
}

type MulticastConfiguration struct {
	Address   IPAddress
	Port      int
	TTL       int
	AutoStart bool
}

type IPAddress struct {
	Type        string
	IPv4Address string
	IPv6Address string
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       string
	Bitrate        int
	SampleRate     int
	Multicast      MulticastConfiguration
	SessionTimeout time.Duration
}

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration
	RuleEngineConfiguration      RuleEngineConfiguration
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config
	Extension       map[string]string
}

type RuleEngineConfiguration struct {
	Rule      Config
	Extension map[string]string
}

type Config struct {
	Name       string
	Type       QName
	Parameters ItemList
}

type QName struct {
	Namespace string
	LocalPart string
}

type ItemList struct {
	SimpleItem  SimpleItem
	ElementItem ElementItem
	Extension   map[string]string
}

type SimpleItem struct {
	Name  string
	Value string
}

type ElementItem struct {
	Name string
}

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int
	PresetRamp                             int
	PresetTourRamp                         int
	NodeToken                              onvif.ReferenceToken
	DefaultAbsolutePantTiltPositionSpace   string
	DefaultAbsoluteZoomPositionSpace       string
	DefaultRelativePanTiltTranslationSpace string
	DefaultRelativeZoomTranslationSpace    string
	DefaultContinuousPanTiltVelocitySpace  string
	DefaultContinuousZoomVelocitySpace     string
	DefaultPTZSpeed                        PTZSpeed
	DefaultPTZTimeout                      time.Duration
	PanTiltLimits                          PanTiltLimits
	ZoomLimits                             ZoomLimits
	Extension                              PTZConfigurationExtension
}

type PTZSpeed struct {
	PanTilt Vector2D
	Zoom    Vector1D
}

type Vector2D struct {
	X     float64
	Y     float64
	Space string
}

type Vector1D struct {
	X     float64
	Space string
}

type PanTiltLimits struct {
	Range Space2DDescription
}

type Space2DDescription struct {
	URI    string
	XRange FloatRange
	YRange FloatRange
}

type ZoomLimits struct {
	Range Space1DDescription
}

type Space1DDescription struct {
	URI    string
	XRange FloatRange
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection
	Extension          map[string]string
}

type PTControlDirection struct {
	EFlip     EFlip
	Reverse   Reverse
	Extension map[string]string
}

type EFlip struct {
	Mode EFlipMode
}

type EFlipMode string

type Reverse struct {
	Mode string
}

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string
	PTZStatus                    PTZFilter
	Events                       EventSubscription
	Analytics                    bool
	Multicast                    MulticastConfiguration
	SessionTimeout               time.Duration
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration
	Extension                    map[string]string
}

type PTZFilter struct {
	Status   bool
	Position bool
}

type EventSubscription struct {
	Filter             map[string]string
	SubscriptionPolicy map[string]string
}

type ProfileExtension struct {
	AudioOutputConfiguration  AudioOutputConfiguration
	AudioDecoderConfiguration AudioDecoderConfiguration
	Extension                 map[string]string
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken onvif.ReferenceToken
	SendPrimacy string
	OutputLevel int
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles    int
	BoundsRange                IntRectangleRange
	VideoSourceTokensAvailable onvif.ReferenceToken
	Extension                  VideoSourceConfigurationOptionsExtension
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    RotateOptions
	Extension VideoSourceConfigurationOptionsExtension2
}

type RotateOptions struct {
	Mode       string
	DegreeList IntList
	Extension  map[string]string
}

type IntList struct {
	Items []int
}

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode string
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	JPEG         JpegOptions
	MPEG4        Mpeg4Options
	H264         H264Options
	Extension    VideoEncoderOptionsExtension
}

type JpegOptions struct {
	ResolutionsAvailable  VideoResolution
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
}

type Mpeg4Options struct {
	ResolutionsAvailable   VideoResolution
	GovLengthRange         IntRange
	FrameRateRange         IntRange
	EncodingIntervalRange  IntRange
	Mpeg4ProfilesSupported string
}

type H264Options struct {
	ResolutionsAvailable  VideoResolution
	GovLengthRange        IntRange
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
	H264ProfilesSupported string
}

type VideoEncoderOptionsExtension struct {
	JPEG      JpegOptions2
	MPEG4     Mpeg4Options2
	H264      H264Options2
	Extension map[string]string
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange
}

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable onvif.ReferenceToken
	Extension            map[string]string
}

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption
}

type AudioEncoderConfigurationOption struct {
	Encoding       string
	BitrateList    IntList
	SampleRateList IntList
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions
	Extension              MetadataConfigurationOptionsExtension
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   bool
	ZoomStatusSupported      bool
	PanTiltPositionSupported bool
	ZoomPositionSupported    bool
	Extension                map[string]string
}

type MetadataConfigurationOptionsExtension struct {
	CompressionType string
	Extension       map[string]string
}

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable onvif.ReferenceToken
	SendPrimacyOptions    string
	OutputLevelRange      IntRange
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  AACDecOptions
	G711DecOptions G711DecOptions
	G726DecOptions G726DecOptions
	Extension      map[string]string
}

type AACDecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G711DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G726DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type StreamSetup struct {
	Stream    StreamType
	Transport Transport
}

type StreamType string

type Transport struct {
	Protocol TransportProtocol
	Tunnel   *Transport
}

// enum
type TransportProtocol string

type MediaUri struct {
	Uri                 string
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             time.Duration
}

type VideoSourceMode struct {
	Token         onvif.ReferenceToken
	Enabled       bool
	MaxFramerate  float64
	MaxResolution VideoResolution
	Encodings     EncodingTypes
	Reboot        bool
	Description   Description
	Extension     map[string]string
}

type EncodingTypes struct {
	EncodingTypes map[string]string
}

type Description struct {
	Description string
}

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs
	Type                string
	PositionOption      string
	TextOption          OSDTextOptions
	ImageOption         OSDImgOptions
	Extension           map[string]string
}

type MaximumNumberOfOSDs struct {
	Total       int
	Image       int
	PlainText   int
	Date        int
	Time        int
	DateAndTime int
}

type OSDTextOptions struct {
	Type            string
	FontSizeRange   IntRange
	DateFormat      string
	TimeFormat      string
	FontColor       OSDColorOptions
	BackgroundColor OSDColorOptions
	Extension       map[string]string
}

type OSDColorOptions struct {
	Color       ColorOptions
	Transparent IntRange
	Extension   map[string]string
}

type ColorOptions struct {
	ColorList       Color
	ColorspaceRange ColorspaceRange
}

type ColorspaceRange struct {
	X          FloatRange
	Y          FloatRange
	Z          FloatRange
	Colorspace string
}

type OSDImgOptions struct {
	FormatsSupported StringAttrList
	MaxSize          int
	MaxWidth         int
	MaxHeight        int

	ImagePath string
	Extension map[string]string
}

type StringAttrList struct {
	AttrList map[string]string
}

//PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      bool
	GeoMove                bool
	Name                   string
	SupportedPTZSpaces     PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported          bool
	AuxiliaryCommands      map[string]string
	Extension              PTZNodeExtension
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    Space2DDescription
	AbsoluteZoomPositionSpace       Space1DDescription
	RelativePanTiltTranslationSpace Space2DDescription
	RelativeZoomTranslationSpace    Space1DDescription
	ContinuousPanTiltVelocitySpace  Space2DDescription
	ContinuousZoomVelocitySpace     Space1DDescription
	PanTiltSpeedSpace               Space1DDescription
	ZoomSpeedSpace                  Space1DDescription
	Extension                       map[string]string
}

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported
	Extension           map[string]string
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int
	PTZPresetTourOperation     string
	Extension                  map[string]string
}

type PTZConfigurationOptions struct {
	PTZRamps           IntAttrList
	Spaces             PTZSpaces
	PTZTimeout         DurationRange
	PTControlDirection PTControlDirectionOptions
	Extension          map[string]string
}

type IntAttrList struct {
	IntAttrList []int
}

type DurationRange struct {
	Min time.Duration
	Max time.Duration
}

type PTControlDirectionOptions struct {
	EFlip     EFlipOptions
	Reverse   ReverseOptions
	Extension map[string]string
}

type EFlipOptions struct {
	Mode      EFlipMode
	Extension map[string]string
}

type ReverseOptions struct {
	Mode      string
	Extension map[string]string
}

type PTZPreset struct {
	Token       onvif.ReferenceToken
	Name        string
	PTZPosition PTZVector
}

type PTZVector struct {
	PanTilt Vector2D
	Zoom    Vector1D
}

type PTZStatus struct {
	Position   PTZVector
	MoveStatus PTZMoveStatus
	Error      string
	UtcTime    time.Time
}

type PTZMoveStatus struct {
	PanTilt MoveStatus
	Zoom    MoveStatus
}

type MoveStatus struct {
	Status string
}

type GeoLocation struct {
	Lon       float64
	Lat       float64
	Elevation float32
}

type PresetTour struct {
	Token             onvif.ReferenceToken
	Name              string
	Status            PTZPresetTourStatus
	AutoStart         bool
	StartingCondition PTZPresetTourStartingCondition
	TourSpot          PTZPresetTourSpot
	Extension         map[string]string
}

type PTZPresetTourStatus struct {
	State           string
	CurrentTourSpot PTZPresetTourSpot
	Extension       map[string]string
}

type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail
	Speed        PTZSpeed
	StayTime     time.Duration
	Extension    map[string]string
}

type PTZPresetTourPresetDetail struct {
	PresetToken   onvif.ReferenceToken
	Home          bool
	PTZPosition   PTZVector
	TypeExtension map[string]string
}

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder bool
	RecurringTime     int
	RecurringDuration time.Duration
	Direction         string
	Extension         map[string]string
}

type PTZPresetTourOptions struct {
	AutoStart         bool
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot          PTZPresetTourSpotOptions
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     IntRange
	RecurringDuration DurationRange
	Direction         string
	Extension         map[string]string
}

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions
	StayTime     DurationRange
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          onvif.ReferenceToken
	Home                 bool
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace    Space1DDescription
	Extension            map[string]string
}

// Capabilities of device
type Capabilities struct {
	Analytics AnalyticsCapabilities
	Device    DeviceCapabilities
	Events    EventCapabilities
	Imaging   ImagingCapabilities
	Media     MediaCapabilities
	PTZ       PTZCapabilities
	Extension CapabilitiesExtension
}

// AnalyticsCapabilities Check
type AnalyticsCapabilities struct {
	XAddr                  string
	RuleSupport            bool
	AnalyticsModuleSupport bool
}

// DeviceCapabilities Check
type DeviceCapabilities struct {
	XAddr     string
	Network   NetworkCapabilities
	System    SystemCapabilities
	IO        IOCapabilities
	Security  SecurityCapabilities
	Extension map[string]string
}

// NetworkCapabilities Check
type NetworkCapabilities struct {
	IPFilter          bool
	ZeroConfiguration bool
	IPVersion6        bool
	DynDNS            bool
	Extension         NetworkCapabilitiesExtension
}

// NetworkCapabilitiesExtension Check
type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool
	Extension          map[string]string
}

// SystemCapabilities check
type SystemCapabilities struct {
	DiscoveryResolve  bool
	DiscoveryBye      bool
	RemoteDiscovery   bool
	SystemBackup      bool
	SystemLogging     bool
	FirmwareUpgrade   bool
	SupportedVersions OnvifVersion
	Extension         map[string]string
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    bool
	HttpSystemBackup       bool
	HttpSystemLogging      bool
	HttpSupportInformation bool
	Extension              map[string]string
}

type IOCapabilities struct {
	InputConnectors int
	RelayOutputs    int
	Extension       IOCapabilitiesExtension
}

type IOCapabilitiesExtension struct {
	Auxiliary         bool
	AuxiliaryCommands string
	Extension         map[string]string
}

type SecurityCapabilities struct {
	TLS1_1               bool
	TLS1_2               bool
	OnboardKeyGeneration bool
	AccessPolicyConfig   bool
	X_509Token           bool
	SAMLToken            bool
	KerberosToken        bool
	RELToken             bool
	Extension            SecurityCapabilitiesExtension
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    bool
	Extension SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              bool
	SupportedEAPMethod int
	RemoteUserHandling bool
}

type EventCapabilities struct {
	XAddr                                         string
	WSSubscriptionPolicySupport                   bool
	WSPullPointSupport                            bool
	WSPausableSubscriptionManagerInterfaceSupport bool
}

type ImagingCapabilities struct {
	XAddr string
}

type MediaCapabilities struct {
	XAddr                 string
	StreamingCapabilities RealTimeStreamingCapabilities
	Extension             MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast bool
	RTP_TCP      bool
	RTP_RTSP_TCP bool
	Extension    map[string]string
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int
}

type PTZCapabilities struct {
	XAddr string
}

type CapabilitiesExtension struct {
	DeviceIO        DeviceIOCapabilities
	Display         DisplayCapabilities
	Recording       RecordingCapabilities
	Search          SearchCapabilities
	Replay          ReplayCapabilities
	Receiver        ReceiverCapabilities
	AnalyticsDevice AnalyticsDeviceCapabilities
	Extensions      map[string]string
}

type DeviceIOCapabilities struct {
	XAddr        string
	VideoSources int
	VideoOutputs int
	AudioSources int
	AudioOutputs int
	RelayOutputs int
}

type DisplayCapabilities struct {
	XAddr       string
	FixedLayout bool
}

type RecordingCapabilities struct {
	XAddr              string
	ReceiverSource     bool
	MediaProfileSource bool
	DynamicRecordings  bool
	DynamicTracks      bool
	MaxStringLength    int
}

type SearchCapabilities struct {
	XAddr          string
	MetadataSearch bool
}

type ReplayCapabilities struct {
	XAddr string
}

type ReceiverCapabilities struct {
	XAddr                string
	RTP_Multicast        bool
	RTP_TCP              bool
	RTP_RTSP_TCP         bool
	SupportedReceivers   int
	MaximumRTSPURILength int
}

type AnalyticsDeviceCapabilities struct {
	XAddr       string
	RuleSupport bool
	Extension   map[string]string
}

type HostnameInformation struct {
	FromDHCP  bool
	Name      string
	Extension map[string]string
}

type DNSInformation struct {
	FromDHCP     bool
	SearchDomain string
	DNSFromDHCP  IPAddress
	DNSManual    IPAddress
	Extension    map[string]string
}

type NTPInformation struct {
	FromDHCP    bool
	NTPFromDHCP string
	NTPManual   string
	Extension   map[string]string
}

type DynamicDNSInformation struct {
	Type      DynamicDNSType
	Name      string
	TTL       time.Duration
	Extension map[string]string
}

// TODO: enumeration
type DynamicDNSType string

type NetworkInterface struct {
	DeviceEntity
	Enabled   bool
	Info      NetworkInterfaceInfo
	Link      NetworkInterfaceLink
	IPv4      IPv4NetworkInterface
	IPv6      IPv6NetworkInterface
	Extension NetworkInterfaceExtension
}

type NetworkInterfaceInfo struct {
	Name      string
	HwAddress string
	MTU       int
}

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting
	OperSettings  NetworkInterfaceConnectionSetting
	InterfaceType IANA_IfTypes
}

type IANA_IfTypes int

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation bool
	Speed           int
	Duplex          Duplex
}

// TODO: enum
type Duplex string

type NetworkInterfaceExtension struct {
	InterfaceType IANA_IfTypes
	Dot3          string
	Dot11         Dot11Configuration
	Extension     map[string]string
}

type Dot11Configuration struct {
	SSID     string
	Mode     Dot11StationMode
	Alias    string
	Priority int64
	Security Dot11SecurityConfiguration
}

type Dot11SecurityConfiguration struct {
	Mode      Dot11SecurityMode
	Algorithm Dot11Cipher
	PSK       Dot11PSKSet
	Dot1X     onvif.ReferenceToken
	Extension map[string]string
}

type Dot11PSKSet struct {
	Key        string
	Passphrase string
	Extension  map[string]string
}

// TODO: enumeration
type Dot11Cipher string

// TODO: enumeration
type Dot11SecurityMode string

// TODO: enumeration
type Dot11StationMode string

type IPv6NetworkInterface struct {
	Enabled bool
	Config  IPv6Configuration
}

type IPv6Configuration struct {
	AcceptRouterAdvert bool
	DHCP               IPv6DHCPConfiguration
	Manual             PrefixedIPv6Address
	LinkLocal          PrefixedIPv6Address
	FromDHCP           PrefixedIPv6Address
	FromRA             PrefixedIPv6Address
	Extension          map[string]string
}

type PrefixedIPv6Address struct {
	Address      string
	PrefixLength int
}

// TODO: enumeration
type IPv6DHCPConfiguration string

type IPv4NetworkInterface struct {
	Enabled bool
	Config  IPv4Configuration
}

type IPv4Configuration struct {
	Manual    PrefixedIPv4Address
	LinkLocal PrefixedIPv4Address
	FromDHCP  PrefixedIPv4Address
	DHCP      bool
}

// optional, unbounded
type PrefixedIPv4Address struct {
	Address      string
	PrefixLength int
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   bool
	Link      NetworkInterfaceConnectionSetting
	MTU       int
	IPv4      IPv4NetworkInterfaceSetConfiguration
	IPv6      IPv6NetworkInterfaceSetConfiguration
	Extension NetworkInterfaceSetConfigurationExtension
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      string
	Dot11     Dot11Configuration
	Extension map[string]string
}

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            bool
	AcceptRouterAdvert bool
	Manual             PrefixedIPv6Address
	DHCP               IPv6DHCPConfiguration
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled bool
	Manual  PrefixedIPv4Address
	DHCP    bool
}

type NetworkProtocol struct {
	Name      string
	Enabled   bool
	Port      int
	Extension map[string]string
}

type NetworkGateway struct {
	IPv4Address string
	IPv6Address string
}

type NetworkZeroConfiguration struct {
	InterfaceToken onvif.ReferenceToken
	Enabled        bool
	Addresses      string
	Extension      NetworkZeroConfigurationExtension
}

type NetworkZeroConfigurationExtension struct {
	Additional *NetworkZeroConfiguration
	Extension  map[string]string
}

type IPAddressFilter struct {
	Type        string
	IPv4Address PrefixedIPv4Address
	IPv6Address PrefixedIPv6Address
	Extension   map[string]string
}

type BinaryData struct {
	X    string
	Data []byte
}

type Certificate struct {
	CertificateID string
	Certificate   BinaryData
}

type CertificateStatus struct {
	CertificateID string
	Status        bool
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings
}

type RelayOutputSettings struct {
	Mode      RelayMode
	DelayTime time.Duration
	IdleState RelayIdleState
}

type RelayIdleState string

type RelayMode string

type RelayLogicalState string

type CertificateWithPrivateKey struct {
	CertificateID string
	Certificate   BinaryData
	PrivateKey    BinaryData
}

type CertificateInformation struct {
	CertificateID      string
	IssuerDN           string
	SubjectDN          string
	KeyUsage           CertificateUsage
	ExtendedKeyUsage   CertificateUsage
	KeyLength          int
	Version            string
	SerialNum          string
	SignatureAlgorithm string
	Validity           DateTimeRange
	Extension          map[string]string
}

type DateTimeRange struct {
	From  time.Time
	Until time.Time
}

type CertificateUsage struct {
	Critical         bool
	CertificateUsage string
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken onvif.ReferenceToken
	Identity                string
	AnonymousID             string
	EAPMethod               int
	CACertificateID         string
	EAPMethodConfiguration  EAPMethodConfiguration
	Extension               map[string]string
}

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration
	Password         string
	Extension        map[string]string
}

type TLSConfiguration struct {
	CertificateID string
}

type Dot11Capabilities struct {
	TKIP                  bool
	ScanAvailableNetworks bool
	MultipleConfiguration bool
	AdHocStationMode      bool
	WEP                   bool
}

type Dot11Status struct {
	SSID              string
	BSSID             string
	PairCipher        Dot11Cipher
	GroupCipher       Dot11Cipher
	SignalStrength    string
	ActiveConfigAlias onvif.ReferenceToken
}

type Dot11AvailableNetworks struct {
	SSID                  string
	BSSID                 string
	AuthAndMangementSuite string
	PairCipher            Dot11Cipher
	GroupCipher           Dot11Cipher
	SignalStrength        string
	Extension             map[string]string
}

type SystemLogUriList struct {
	SystemLog SystemLogUri
}

type SystemLogUri struct {
	Type string
	Uri  string
}

type LocationEntity struct {
	Entity    string
	Token     onvif.ReferenceToken
	Fixed     bool
	GeoSource string
	AutoGeo   bool

	GeoLocation      GeoLocation
	GeoOrientation   GeoOrientation
	LocalLocation    LocalLocation
	LocalOrientation LocalOrientation
}

type LocalOrientation struct {
	Lon       float64
	Lat       float64
	Elevation float32
}

type LocalLocation struct {
	X float32
	Y float32
	Z float32
}

type GeoOrientation struct {
	Roll  float32
	Pitch float32
	Yaw   float32
}

type FocusMove struct {
	Absolute   AbsoluteFocus
	Relative   RelativeFocus
	Continuous ContinuousFocus
}

type ContinuousFocus struct {
	Speed float32
}

type RelativeFocus struct {
	Distance float32
	Speed    float32
}

type AbsoluteFocus struct {
	Position float32
	Speed    float32
}

type DateTime struct {
	Time Time
	Date Date
}

type Time struct {
	Hour   int
	Minute int
	Second int
}

type Date struct {
	Year  int
	Month int
	Day   int
}
