package device

import (
	"time"

	"github.com/jjbarbosa7/onvif/xsd/onvif"
)

// Translation of onvif types from wsdl to go

type TimeDuration string

type DeviceEntity struct {
	Token onvif.ReferenceToken `bson:"token" json:"token"`
}

type OnvifVersion struct {
	Major int `bson:"major" json:"major"`
	Minor int `bson:"minor" json:"minor"`
}

type SystemDateTime struct {
	DateTimeType    string            `bson:"dateTimeType" json:"dateTimeType"`
	DaylightSavings bool              `bson:"daylightSavings" json:"daylightSavings"`
	TimeZone        string            `bson:"timeZone" json:"timeZone"`
	UTCDateTime     time.Time         `bson:"utcDateTime" json:"utcDateTime"`
	LocalDateTime   time.Time         `bson:"localDateTime" json:"localDateTime"`
	Extension       map[string]string `bson:"extension" json:"extension"`
}

type IntRectangle struct {
	X      int `bson:"x" json:"x"`
	Y      int `bson:"y" json:"y"`
	Width  int `bson:"width" json:"width"`
	Height int `bson:"height" json:"height"`
}

type IntRectangleRange struct {
	XRange      IntRange `bson:"xRange" json:"xRange"`
	YRange      IntRange `bson:"yRange" json:"yRange"`
	WidthRange  IntRange `bson:"widthRange" json:"widthRange"`
	HeightRange IntRange `bson:"heightRange" json:"heightRange"`
}

type IntRange struct {
	Min int `bson:"min" json:"min"`
	Max int `bson:"max" json:"max"`
}

type FloatRange struct {
	Min float64 `bson:"min" json:"min"`
	Max float64 `bson:"max" json:"max"`
}

type OSDConfiguration struct {
	DeviceEntity
	VideoSourceConfigurationToken onvif.ReferenceToken `bson:"videoSourceConfigurationToken" json:"videoSourceConfigurationToken"`
	Type                          string               `bson:"type" json:"type"`
	Position                      OSDPosConfiguration  `bson:"position" json:"position"`
	TextString                    OSDTextConfiguration `bson:"textString" json:"textString"`
	Image                         OSDImgConfiguration  `bson:"image" json:"image"`
	Extension                     map[string]string    `bson:"extension" json:"extension"`
}

type OSDPosConfiguration struct {
	Type      string            `bson:"type" json:"type"`
	Pos       Vector            `bson:"pos" json:"pos"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type Vector struct {
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
}

type OSDTextConfiguration struct {
	IsPersistentText bool              `bson:"isPersistentText" json:"isPersistentText"`
	Type             string            `bson:"type" json:"type"`
	DateFormat       string            `bson:"dateFormat" json:"dateFormat"`
	TimeFormat       string            `bson:"timeFormat" json:"timeFormat"`
	FontSize         int               `bson:"fontSize" json:"fontSize"`
	FontColor        OSDColor          `bson:"fontColor" json:"fontColor"`
	BackgroundColor  OSDColor          `bson:"backgroundColor" json:"backgroundColor"`
	PlainText        string            `bson:"plainText" json:"plainText"`
	Extension        map[string]string `bson:"extension" json:"extension"`
}

type OSDColor struct {
	Transparent int   `bson:"transparent" json:"transparent"`
	Color       Color `bson:"color" json:"color"`
}

type Color struct {
	X          float64 `bson:"x" json:"x"`
	Y          float64 `bson:"y" json:"y"`
	Z          float64 `bson:"z" json:"z"`
	Colorspace string  `bson:"colorspace" json:"colorspace"`
}

type OSDImgConfiguration struct {
	ImgPath   string            `bson:"imgPath" json:"imgPath"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type VideoSource struct {
	DeviceEntity
	Framerate  float64              `bson:"framerate" json:"framerate"`
	Resolution VideoResolution      `bson:"resolution" json:"resolution"`
	Imaging    ImagingSettings      `bson:"imaging" json:"imaging"`
	Extension  VideoSourceExtension `bson:"extension" json:"extension"`
}

type VideoResolution struct {
	Width  int `bson:"width" json:"width"`
	Height int `bson:"height" json:"height"`
}

type ImagingSettings struct {
	BacklightCompensation BacklightCompensation `bson:"backlightCompensation" json:"backlightCompensation"`
	Brightness            float64               `bson:"brightness" json:"brightness"`
	ColorSaturation       float64               `bson:"colorSaturation" json:"colorSaturation"`
	Contrast              float64               `bson:"contrast" json:"contrast"`
	Exposure              Exposure              `bson:"exposure" json:"exposure"`
	Focus                 FocusConfiguration    `bson:"focus" json:"focus"`
	IrCutFilter           string                `bson:"irCutFilter" json:"irCutFilter"`
	Sharpness             float64               `bson:"sharpness" json:"sharpness"`
	WideDynamicRange      WideDynamicRange      `bson:"wideDynamicRange" json:"wideDynamicRange"`
	WhiteBalance          WhiteBalance          `bson:"whiteBalance" json:"whiteBalance"`
	Extension             map[string]string     `bson:"extension" json:"extension"`
}

type BacklightCompensation struct {
	Mode  string  `bson:"mode" json:"mode"`
	Level float64 `bson:"level" json:"level"`
}

type Exposure struct {
	Mode            string    `bson:"mode" json:"mode"`
	Priority        string    `bson:"priority" json:"priority"`
	Window          Rectangle `bson:"window" json:"window"`
	MinExposureTime float64   `bson:"minExposureTime" json:"minExposureTime"`
	MaxExposureTime float64   `bson:"maxExposureTime" json:"maxExposureTime"`
	MinGain         float64   `bson:"minGain" json:"minGain"`
	MaxGain         float64   `bson:"maxGain" json:"maxGain"`
	MinIris         float64   `bson:"minIris" json:"minIris"`
	MaxIris         float64   `bson:"maxIris" json:"maxIris"`
	ExposureTime    float64   `bson:"exposureTime" json:"exposureTime"`
	Gain            float64   `bson:"gain" json:"gain"`
	Iris            float64   `bson:"iris" json:"iris"`
}

type Rectangle struct {
	Bottom float64 `bson:"bottom" json:"bottom"`
	Top    float64 `bson:"top" json:"top"`
	Right  float64 `bson:"right" json:"right"`
	Left   float64 `bson:"left" json:"left"`
}

type FocusConfiguration struct {
	AutoFocusMode string  `bson:"autoFocusMode" json:"autoFocusMode"`
	DefaultSpeed  float64 `bson:"defaultSpeed" json:"defaultSpeed"`
	NearLimit     float64 `bson:"nearLimit" json:"nearLimit"`
	FarLimit      float64 `bson:"farLimit" json:"farLimit"`
}

type WideDynamicRange struct {
	Mode  string  `bson:"mode" json:"mode"`
	Level float64 `bson:"level" json:"level"`
}

type WhiteBalance struct {
	Mode   string  `bson:"mode" json:"mode"`
	CrGain float64 `bson:"crGain" json:"crGain"`
	CbGain float64 `bson:"cbGain" json:"cbGain"`
}

type VideoSourceExtension struct {
	Imaging   ImagingSettings20 `bson:"imaging" json:"imaging"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type ImagingSettings20 struct {
	BacklightCompensation *BacklightCompensation20    `bson:"backlightCompensation" json:"backlightCompensation"`
	Brightness            float64                     `bson:"brightness" json:"brightness"`
	ColorSaturation       float64                     `bson:"colorSaturation" json:"colorSaturation"`
	Contrast              float64                     `bson:"contrast" json:"contrast"`
	Exposure              *Exposure20                 `bson:"exposure" json:"exposure"`
	Focus                 *FocusConfiguration20       `bson:"focus" json:"focus"`
	IrCutFilter           *string                     `bson:"irCutFilter" json:"irCutFilter"`
	Sharpness             float64                     `bson:"sharpness" json:"sharpness"`
	WideDynamicRange      *WideDynamicRange20         `bson:"wideDynamicRange" json:"wideDynamicRange"`
	WhiteBalance          *WhiteBalance20             `bson:"whiteBalance" json:"whiteBalance"`
	Extension             *ImagingSettingsExtension20 `bson:"extension" json:"extension"`
}
type BacklightCompensation20 struct {
	Mode  string  `bson:"mode" json:"mode"`
	Level float64 `bson:"level" json:"level"`
}

type Exposure20 struct {
	Mode            string    `bson:"mode" json:"mode"`
	Priority        string    `bson:"priority" json:"priority"`
	Window          Rectangle `bson:"window" json:"window"`
	MinExposureTime float64   `bson:"minExposureTime" json:"minExposureTime"`
	MaxExposureTime float64   `bson:"maxExposureTime" json:"maxExposureTime"`
	MinGain         float64   `bson:"minGain" json:"minGain"`
	MaxGain         float64   `bson:"maxGain" json:"maxGain"`
	MinIris         float64   `bson:"minIris" json:"minIris"`
	MaxIris         float64   `bson:"maxIris" json:"maxIris"`
	ExposureTime    float64   `bson:"exposureTime" json:"exposureTime"`
	Gain            float64   `bson:"gain" json:"gain"`
	Iris            float64   `bson:"iris" json:"iris"`
}

type FocusConfiguration20 struct {
	AutoFocusMode string            `bson:"autoFocusMode" json:"autoFocusMode"`
	DefaultSpeed  float64           `bson:"defaultSpeed" json:"defaultSpeed"`
	NearLimit     float64           `bson:"nearLimit" json:"nearLimit"`
	FarLimit      float64           `bson:"farLimit" json:"farLimit"`
	Extension     map[string]string `bson:"extension" json:"extension"`
}

type WideDynamicRange20 struct {
	Mode  string  `bson:"mode" json:"mode"`
	Level float64 `bson:"level" json:"level"`
}

type WhiteBalance20 struct {
	Mode      string            `bson:"mode" json:"mode"`
	CrGain    float64           `bson:"crGain" json:"crGain"`
	CbGain    float64           `bson:"cbGain" json:"cbGain"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization          `bson:"imageStabilization" json:"imageStabilization"`
	Extension          ImagingSettingsExtension202 `bson:"extension" json:"extension"`
}

type ImageStabilization struct {
	Mode      string            `bson:"mode" json:"mode"`
	Level     float64           `bson:"level" json:"level"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment   `bson:"irCutFilterAutoAdjustment" json:"irCutFilterAutoAdjustment"`
	Extension                 ImagingSettingsExtension203 `bson:"extension" json:"extension"`
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string            `bson:"boundaryType" json:"boundaryType"`
	BoundaryOffset float64           `bson:"boundaryOffset" json:"boundaryOffset"`
	ResponseTime   TimeDuration      `bson:"responseTime" json:"responseTime"`
	Extension      map[string]string `bson:"extension" json:"extension"`
}

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation  `bson:"toneCompensation" json:"toneCompensation"`
	Defogging        Defogging         `bson:"defogging" json:"defogging"`
	NoiseReduction   NoiseReduction    `bson:"noiseReduction" json:"noiseReduction"`
	Extension        map[string]string `bson:"extension" json:"extension"`
}

type ToneCompensation struct {
	Mode      string            `bson:"mode" json:"mode"`
	Level     float64           `bson:"level" json:"level"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type Defogging struct {
	Mode      string            `bson:"mode" json:"mode"`
	Level     float64           `bson:"level" json:"level"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type NoiseReduction struct {
	Level float64 `bson:"level" json:"level"`
}

type AudioSource struct {
	DeviceEntity
	Channels int `bson:"channels" json:"channels"`
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token                       onvif.ReferenceToken        `bson:"token" json:"token"`
	Fixed                       bool                        `bson:"fixed" json:"fixed"`
	Name                        string                      `bson:"name" json:"name"`
	VideoSourceConfiguration    VideoSourceConfiguration    `bson:"videoSourceConfiguration" json:"videoSourceConfiguration"`
	AudioSourceConfiguration    AudioSourceConfiguration    `bson:"audioSourceConfiguration" json:"audioSourceConfiguration"`
	VideoEncoderConfiguration   VideoEncoderConfiguration   `bson:"videoEncoderConfiguration" json:"videoEncoderConfiguration"`
	AudioEncoderConfiguration   AudioEncoderConfiguration   `bson:"audioEncoderConfiguration" json:"audioEncoderConfiguration"`
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `bson:"videoAnalyticsConfiguration" json:"videoAnalyticsConfiguration"`
	PTZConfiguration            PTZConfiguration            `bson:"ptzConfiguration" json:"ptzConfiguration"`
	MetadataConfiguration       MetadataConfiguration       `bson:"metadataConfiguration" json:"metadataConfiguration"`
	Extension                   ProfileExtension            `bson:"extension" json:"extension"`
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode    string                            `bson:"viewMode" json:"viewMode"`
	SourceToken onvif.ReferenceToken              `bson:"sourceToken" json:"sourceToken"`
	Bounds      IntRectangle                      `bson:"bounds" json:"bounds"`
	Extension   VideoSourceConfigurationExtension `bson:"extension" json:"extension"`
}

type ConfigurationEntity struct {
	Token    onvif.ReferenceToken `bson:"token" json:"token"`
	Name     string               `bson:"name" json:"name"`
	UseCount int                  `bson:"useCount" json:"useCount"`
}

type VideoSourceConfigurationExtension struct {
	Rotate    Rotate                             `bson:"rotate" json:"rotate"`
	Extension VideoSourceConfigurationExtension2 `bson:"extension" json:"extension"`
}

type Rotate struct {
	Mode      string            `bson:"mode" json:"mode"`
	Degree    int               `bson:"degree" json:"degree"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription  LensDescription  `bson:"lensDescription" json:"lensDescription"`
	SceneOrientation SceneOrientation `bson:"sceneOrientation" json:"sceneOrientation"`
}

type LensDescription struct {
	FocalLength float64        `bson:"focalLength" json:"focalLength"`
	Offset      LensOffset     `bson:"offset" json:"offset"`
	Projection  LensProjection `bson:"projection" json:"projection"`
	XFactor     float64        `bson:"xFactor" json:"xFactor"`
}

type LensOffset struct {
	X float64 `bson:"x" json:"x"`
	Y float64 `bson:"y" json:"y"`
}

type LensProjection struct {
	Angle         float64 `bson:"angle" json:"angle"`
	Radius        float64 `bson:"radius" json:"radius"`
	Transmittance float64 `bson:"transmittance" json:"transmittance"`
}

type SceneOrientation struct {
	Mode        string `bson:"mode" json:"mode"`
	Orientation string `bson:"orientation" json:"orientation"`
}

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken onvif.ReferenceToken `bson:"sourceToken" json:"sourceToken"`
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       string                 `bson:"encoding" json:"encoding"`
	Resolution     VideoResolution        `bson:"resolution" json:"resolution"`
	Quality        float64                `bson:"quality" json:"quality"`
	RateControl    VideoRateControl       `bson:"rateControl" json:"rateControl"`
	MPEG4          Mpeg4Configuration     `bson:"mpeg4" json:"mpeg4"`
	H264           H264Configuration      `bson:"h264" json:"h264"`
	Multicast      MulticastConfiguration `bson:"multicast" json:"multicast"`
	SessionTimeout TimeDuration           `bson:"sessionTimeout" json:"sessionTimeout"`
}

type VideoRateControl struct {
	FrameRateLimit   int `bson:"frameRateLimit" json:"frameRateLimit"`
	EncodingInterval int `bson:"encodingInterval" json:"encodingInterval"`
	BitrateLimit     int `bson:"bitrateLimit" json:"bitrateLimit"`
}

type Mpeg4Configuration struct {
	GovLength    int    `bson:"govLength" json:"govLength"`
	Mpeg4Profile string `bson:"mpeg4Profile" json:"mpeg4Profile"`
}

type H264Configuration struct {
	GovLength   int    `bson:"govLength" json:"govLength"`
	H264Profile string `bson:"h264Profile" json:"h264Profile"`
}

type MulticastConfiguration struct {
	Address   IPAddress `bson:"address" json:"address"`
	Port      int       `bson:"port" json:"port"`
	TTL       int       `bson:"ttl" json:"ttl"`
	AutoStart bool      `bson:"autoStart" json:"autoStart"`
}

type IPAddress struct {
	Type        string `bson:"type" json:"type"`
	IPv4Address string `bson:"ipv4Address" json:"ipv4Address"`
	IPv6Address string `bson:"ipv6Address" json:"ipv6Address"`
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       string                 `bson:"encoding" json:"encoding"`
	Bitrate        int                    `bson:"bitrate" json:"bitrate"`
	SampleRate     int                    `bson:"sampleRate" json:"sampleRate"`
	Multicast      MulticastConfiguration `bson:"multicast" json:"multicast"`
	SessionTimeout TimeDuration           `bson:"sessionTimeout" json:"sessionTimeout"`
}

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `bson:"analyticsEngineConfiguration" json:"analyticsEngineConfiguration"`
	RuleEngineConfiguration      RuleEngineConfiguration      `bson:"ruleEngineConfiguration" json:"ruleEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config            `bson:"analyticsModule" json:"analyticsModule"`
	Extension       map[string]string `bson:"extension" json:"extension"`
}

type RuleEngineConfiguration struct {
	Rule      Config            `bson:"rule" json:"rule"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type Config struct {
	Name       string   `bson:"name" json:"name"`
	Type       QName    `bson:"type" json:"type"`
	Parameters ItemList `bson:"parameters" json:"parameters"`
}

type QName struct {
	Namespace string `bson:"namespace" json:"namespace"`
	LocalPart string `bson:"localPart" json:"localPart"`
}

type ItemList struct {
	SimpleItem  SimpleItem        `bson:"simpleItem" json:"simpleItem"`
	ElementItem ElementItem       `bson:"elementItem" json:"elementItem"`
	Extension   map[string]string `bson:"extension" json:"extension"`
}

type SimpleItem struct {
	Name  string `bson:"name" json:"name"`
	Value string `bson:"value" json:"value"`
}

type ElementItem struct {
	Name string `bson:"name" json:"name"`
}

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int                       `bson:"moveRamp" json:"moveRamp"`
	PresetRamp                             int                       `bson:"presetRamp" json:"presetRamp"`
	PresetTourRamp                         int                       `bson:"presetTourRamp" json:"presetTourRamp"`
	NodeToken                              onvif.ReferenceToken      `bson:"nodeToken" json:"nodeToken"`
	DefaultAbsolutePantTiltPositionSpace   string                    `bson:"defaultAbsolutePantTiltPositionSpace" json:"defaultAbsolutePantTiltPositionSpace"`
	DefaultAbsoluteZoomPositionSpace       string                    `bson:"defaultAbsoluteZoomPositionSpace" json:"defaultAbsoluteZoomPositionSpace"`
	DefaultRelativePanTiltTranslationSpace string                    `bson:"defaultRelativePanTiltTranslationSpace" json:"defaultRelativePanTiltTranslationSpace"`
	DefaultRelativeZoomTranslationSpace    string                    `bson:"defaultRelativeZoomTranslationSpace" json:"defaultRelativeZoomTranslationSpace"`
	DefaultContinuousPanTiltVelocitySpace  string                    `bson:"defaultContinuousPanTiltVelocitySpace" json:"defaultContinuousPanTiltVelocitySpace"`
	DefaultContinuousZoomVelocitySpace     string                    `bson:"defaultContinuousZoomVelocitySpace" json:"defaultContinuousZoomVelocitySpace"`
	DefaultPTZSpeed                        PTZSpeed                  `bson:"defaultPTZSpeed" json:"defaultPTZSpeed"`
	DefaultPTZTimeout                      TimeDuration              `bson:"defaultPTZTimeout" json:"defaultPTZTimeout"`
	PanTiltLimits                          PanTiltLimits             `bson:"panTiltLimits" json:"panTiltLimits"`
	ZoomLimits                             ZoomLimits                `bson:"zoomLimits" json:"zoomLimits"`
	Extension                              PTZConfigurationExtension `bson:"extension" json:"extension"`
}
type PTZSpeed struct {
	PanTilt Vector2D `bson:"panTilt" json:"panTilt"`
	Zoom    Vector1D `bson:"zoom" json:"zoom"`
}

type Vector2D struct {
	X     float64 `bson:"x" json:"x"`
	Y     float64 `bson:"y" json:"y"`
	Space string  `bson:"space" json:"space"`
}

type Vector1D struct {
	X     float64 `bson:"x" json:"x"`
	Space string  `bson:"space" json:"space"`
}

type PanTiltLimits struct {
	Range Space2DDescription `bson:"range" json:"range"`
}

type Space2DDescription struct {
	URI    string     `bson:"uri" json:"uri"`
	XRange FloatRange `bson:"xRange" json:"xRange"`
	YRange FloatRange `bson:"yRange" json:"yRange"`
}

type ZoomLimits struct {
	Range Space1DDescription `bson:"range" json:"range"`
}

type Space1DDescription struct {
	URI    string     `bson:"uri" json:"uri"`
	XRange FloatRange `bson:"xRange" json:"xRange"`
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection `bson:"ptControlDirection" json:"ptControlDirection"`
	Extension          map[string]string  `bson:"extension" json:"extension"`
}

type PTControlDirection struct {
	EFlip     EFlip             `bson:"eFlip" json:"eFlip"`
	Reverse   Reverse           `bson:"reverse" json:"reverse"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type EFlip struct {
	Mode EFlipMode `bson:"mode" json:"mode"`
}

type EFlipMode string

type Reverse struct {
	Mode string `bson:"mode" json:"mode"`
}

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string                       `bson:"compressionType" json:"compressionType"`
	PTZStatus                    PTZFilter                    `bson:"ptzStatus" json:"ptzStatus"`
	Events                       EventSubscription            `bson:"events" json:"events"`
	Analytics                    bool                         `bson:"analytics" json:"analytics"`
	Multicast                    MulticastConfiguration       `bson:"multicast" json:"multicast"`
	SessionTimeout               TimeDuration                 `bson:"sessionTimeout" json:"sessionTimeout"`
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `bson:"analyticsEngineConfiguration" json:"analyticsEngineConfiguration"`
	Extension                    map[string]string            `bson:"extension" json:"extension"`
}

type PTZFilter struct {
	Status   bool `bson:"status" json:"status"`
	Position bool `bson:"position" json:"position"`
}

type EventSubscription struct {
	Filter             map[string]string `bson:"filter" json:"filter"`
	SubscriptionPolicy map[string]string `bson:"subscriptionPolicy" json:"subscriptionPolicy"`
}

type ProfileExtension struct {
	AudioOutputConfiguration  AudioOutputConfiguration  `bson:"audioOutputConfiguration" json:"audioOutputConfiguration"`
	AudioDecoderConfiguration AudioDecoderConfiguration `bson:"audioDecoderConfiguration" json:"audioDecoderConfiguration"`
	Extension                 map[string]string         `bson:"extension" json:"extension"`
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken onvif.ReferenceToken `bson:"outputToken" json:"outputToken"`
	SendPrimacy string               `bson:"sendPrimacy" json:"sendPrimacy"`
	OutputLevel int                  `bson:"outputLevel" json:"outputLevel"`
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles    int                                      `bson:"maximumNumberOfProfiles" json:"maximumNumberOfProfiles"`
	BoundsRange                IntRectangleRange                        `bson:"boundsRange" json:"boundsRange"`
	VideoSourceTokensAvailable onvif.ReferenceToken                     `bson:"videoSourceTokensAvailable" json:"videoSourceTokensAvailable"`
	Extension                  VideoSourceConfigurationOptionsExtension `bson:"extension" json:"extension"`
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    Rotate                                    `bson:"rotate" json:"rotate"`
	Extension VideoSourceConfigurationOptionsExtension2 `bson:"extension" json:"extension"`
}

type RotateOptions struct {
	Mode       string            `bson:"mode" json:"mode"`
	DegreeList IntList           `bson:"degreeList" json:"degreeList"`
	Extension  map[string]string `bson:"extension" json:"extension"`
}

type IntList struct {
	Items []int `bson:"items" json:"items"`
}

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode string `bson:"sceneOrientationMode" json:"sceneOrientationMode"`
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange                     `bson:"qualityRange" json:"qualityRange"`
	JPEG         JpegOptions                  `bson:"jpeg" json:"jpeg"`
	MPEG4        Mpeg4Options                 `bson:"mpeg4" json:"mpeg4"`
	H264         H264Options                  `bson:"h264" json:"h264"`
	Extension    VideoEncoderOptionsExtension `bson:"extension" json:"extension"`
}

type JpegOptions struct {
	ResolutionsAvailable  VideoResolution `bson:"resolutionsAvailable" json:"resolutionsAvailable"`
	FrameRateRange        IntRange        `bson:"frameRateRange" json:"frameRateRange"`
	EncodingIntervalRange IntRange        `bson:"encodingIntervalRange" json:"encodingIntervalRange"`
}

type Mpeg4Options struct {
	ResolutionsAvailable   VideoResolution `bson:"resolutionsAvailable" json:"resolutionsAvailable"`
	GovLengthRange         IntRange        `bson:"govLengthRange" json:"govLengthRange"`
	FrameRateRange         IntRange        `bson:"frameRateRange" json:"frameRateRange"`
	EncodingIntervalRange  IntRange        `bson:"encodingIntervalRange" json:"encodingIntervalRange"`
	Mpeg4ProfilesSupported string          `bson:"mpeg4ProfilesSupported" json:"mpeg4ProfilesSupported"`
}

type H264Options struct {
	ResolutionsAvailable  VideoResolution `bson:"resolutionsAvailable" json:"resolutionsAvailable"`
	GovLengthRange        IntRange        `bson:"govLengthRange" json:"govLengthRange"`
	FrameRateRange        IntRange        `bson:"frameRateRange" json:"frameRateRange"`
	EncodingIntervalRange IntRange        `bson:"encodingIntervalRange" json:"encodingIntervalRange"`
	H264ProfilesSupported string          `bson:"h264ProfilesSupported" json:"h264ProfilesSupported"`
}

type VideoEncoderOptionsExtension struct {
	JPEG      JpegOptions2      `bson:"jpeg" json:"jpeg"`
	MPEG4     Mpeg4Options2     `bson:"mpeg4" json:"mpeg4"`
	H264      H264Options2      `bson:"h264" json:"h264"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange `bson:"bitrateRange" json:"bitrateRange"`
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange `bson:"bitrateRange" json:"bitrateRange"`
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange `bson:"bitrateRange" json:"bitrateRange"`
}

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable onvif.ReferenceToken `bson:"inputTokensAvailable" json:"inputTokensAvailable"`
	Extension            map[string]string    `bson:"extension" json:"extension"`
}

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption `bson:"options" json:"options"`
}

type AudioEncoderConfigurationOption struct {
	Encoding       string  `bson:"encoding" json:"encoding"`
	BitrateList    IntList `bson:"bitrateList" json:"bitrateList"`
	SampleRateList IntList `bson:"sampleRateList" json:"sampleRateList"`
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions                `bson:"ptzStatusFilterOptions" json:"ptzStatusFilterOptions"`
	Extension              MetadataConfigurationOptionsExtension `bson:"extension" json:"extension"`
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   bool              `bson:"panTiltStatusSupported" json:"panTiltStatusSupported"`
	ZoomStatusSupported      bool              `bson:"zoomStatusSupported" json:"zoomStatusSupported"`
	PanTiltPositionSupported bool              `bson:"panTiltPositionSupported" json:"panTiltPositionSupported"`
	ZoomPositionSupported    bool              `bson:"zoomPositionSupported" json:"zoomPositionSupported"`
	Extension                map[string]string `bson:"extension" json:"extension"`
}

type MetadataConfigurationOptionsExtension struct {
	CompressionType string            `bson:"compressionType" json:"compressionType"`
	Extension       map[string]string `bson:"extension" json:"extension"`
}

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable onvif.ReferenceToken `bson:"outputTokensAvailable" json:"outputTokensAvailable"`
	SendPrimacyOptions    string               `bson:"sendPrimacyOptions" json:"sendPrimacyOptions"`
	OutputLevelRange      IntRange             `bson:"outputLevelRange" json:"outputLevelRange"`
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  AACDecOptions     `bson:"aacDecOptions" json:"aacDecOptions"`
	G711DecOptions G711DecOptions    `bson:"g711DecOptions" json:"g711DecOptions"`
	G726DecOptions G726DecOptions    `bson:"g726DecOptions" json:"g726DecOptions"`
	Extension      map[string]string `bson:"extension" json:"extension"`
}

type AACDecOptions struct {
	Bitrate         IntList `bson:"bitrate" json:"bitrate"`
	SampleRateRange IntList `bson:"sampleRateRange" json:"sampleRateRange"`
}

type G711DecOptions struct {
	Bitrate         IntList `bson:"bitrate" json:"bitrate"`
	SampleRateRange IntList `bson:"sampleRateRange" json:"sampleRateRange"`
}

type G726DecOptions struct {
	Bitrate         IntList `bson:"bitrate" json:"bitrate"`
	SampleRateRange IntList `bson:"sampleRateRange" json:"sampleRateRange"`
}
type StreamSetup struct {
	Stream    string `bson:"stream" json:"stream"`
	Transport string `bson:"transport" json:"transport"`
}

type Transport struct {
	Protocol string     `bson:"protocol" json:"protocol"`
	Tunnel   *Transport `bson:"tunnel" json:"tunnel"`
}

type MediaUri struct {
	Uri                 string       `bson:"uri" json:"uri"`
	InvalidAfterConnect bool         `bson:"invalidAfterConnect" json:"invalidAfterConnect"`
	InvalidAfterReboot  bool         `bson:"invalidAfterReboot" json:"invalidAfterReboot"`
	Timeout             TimeDuration `bson:"timeout" json:"timeout"`
}

type VideoSourceMode struct {
	Token         onvif.ReferenceToken `bson:"token" json:"token"`
	Enabled       bool                 `bson:"enabled" json:"enabled"`
	MaxFramerate  float64              `bson:"maxFramerate" json:"maxFramerate"`
	MaxResolution VideoResolution      `bson:"maxResolution" json:"maxResolution"`
	Encodings     EncodingTypes        `bson:"encodings" json:"encodings"`
	Reboot        bool                 `bson:"reboot" json:"reboot"`
	Description   Description          `bson:"description" json:"description"`
	Extension     map[string]string    `bson:"extension" json:"extension"`
}

type EncodingTypes struct {
	EncodingTypes map[string]string `bson:"encodingTypes" json:"encodingTypes"`
}

type Description struct {
	Description string `bson:"description" json:"description"`
}

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs `bson:"maximumNumberOfOSDs" json:"maximumNumberOfOSDs"`
	Type                string              `bson:"type" json:"type"`
	PositionOption      string              `bson:"positionOption" json:"positionOption"`
	TextOption          OSDTextOptions      `bson:"textOption" json:"textOption"`
	ImageOption         OSDImgOptions       `bson:"imageOption" json:"imageOption"`
	Extension           map[string]string   `bson:"extension" json:"extension"`
}

type MaximumNumberOfOSDs struct {
	Total       int `bson:"total" json:"total"`
	Image       int `bson:"image" json:"image"`
	PlainText   int `bson:"plainText" json:"plainText"`
	Date        int `bson:"date" json:"date"`
	Time        int `bson:"time" json:"time"`
	DateAndTime int `bson:"dateAndTime" json:"dateAndTime"`
}

type OSDTextOptions struct {
	Type            string            `bson:"type" json:"type"`
	FontSizeRange   IntRange          `bson:"fontSizeRange" json:"fontSizeRange"`
	DateFormat      string            `bson:"dateFormat" json:"dateFormat"`
	TimeFormat      string            `bson:"timeFormat" json:"timeFormat"`
	FontColor       OSDColorOptions   `bson:"fontColor" json:"fontColor"`
	BackgroundColor OSDColorOptions   `bson:"backgroundColor" json:"backgroundColor"`
	Extension       map[string]string `bson:"extension" json:"extension"`
}

type OSDColorOptions struct {
	Color       ColorOptions      `bson:"color" json:"color"`
	Transparent IntRange          `bson:"transparent" json:"transparent"`
	Extension   map[string]string `bson:"extension" json:"extension"`
}

type ColorOptions struct {
	ColorList       Color           `bson:"colorList" json:"colorList"`
	ColorspaceRange ColorspaceRange `bson:"colorspaceRange" json:"colorspaceRange"`
}

type ColorspaceRange struct {
	X          FloatRange `bson:"x" json:"x"`
	Y          FloatRange `bson:"y" json:"y"`
	Z          FloatRange `bson:"z" json:"z"`
	Colorspace string     `bson:"colorspace" json:"colorspace"`
}

type OSDImgOptions struct {
	FormatsSupported StringAttrList    `bson:"formatsSupported" json:"formatsSupported"`
	MaxSize          int               `bson:"maxSize" json:"maxSize"`
	MaxWidth         int               `bson:"maxWidth" json:"maxWidth"`
	MaxHeight        int               `bson:"maxHeight" json:"maxHeight"`
	ImagePath        string            `bson:"imagePath" json:"imagePath"`
	Extension        map[string]string `bson:"extension" json:"extension"`
}

type StringAttrList struct {
	AttrList map[string]string `bson:"attrList" json:"attrList"`
}

// PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      bool              `bson:"fixedHomePosition" json:"fixedHomePosition"`
	GeoMove                bool              `bson:"geoMove" json:"geoMove"`
	Name                   string            `bson:"name" json:"name"`
	SupportedPTZSpaces     PTZSpaces         `bson:"supportedPTZSpaces" json:"supportedPTZSpaces"`
	MaximumNumberOfPresets int               `bson:"maximumNumberOfPresets" json:"maximumNumberOfPresets"`
	HomeSupported          bool              `bson:"homeSupported" json:"homeSupported"`
	AuxiliaryCommands      map[string]string `bson:"auxiliaryCommands" json:"auxiliaryCommands"`
	Extension              PTZNodeExtension  `bson:"extension" json:"extension"`
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    Space2DDescription `bson:"absolutePanTiltPositionSpace" json:"absolutePanTiltPositionSpace"`
	AbsoluteZoomPositionSpace       Space1DDescription `bson:"absoluteZoomPositionSpace" json:"absoluteZoomPositionSpace"`
	RelativePanTiltTranslationSpace Space2DDescription `bson:"relativePanTiltTranslationSpace" json:"relativePanTiltTranslationSpace"`
	RelativeZoomTranslationSpace    Space1DDescription `bson:"relativeZoomTranslationSpace" json:"relativeZoomTranslationSpace"`
	ContinuousPanTiltVelocitySpace  Space2DDescription `bson:"continuousPanTiltVelocitySpace" json:"continuousPanTiltVelocitySpace"`
	ContinuousZoomVelocitySpace     Space1DDescription `bson:"continuousZoomVelocitySpace" json:"continuousZoomVelocitySpace"`
	PanTiltSpeedSpace               Space1DDescription `bson:"panTiltSpeedSpace" json:"panTiltSpeedSpace"`
	ZoomSpeedSpace                  Space1DDescription `bson:"zoomSpeedSpace" json:"zoomSpeedSpace"`
	Extension                       map[string]string  `bson:"extension" json:"extension"`
}

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported `bson:"supportedPresetTour" json:"supportedPresetTour"`
	Extension           map[string]string      `bson:"extension" json:"extension"`
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int               `bson:"maximumNumberOfPresetTours" json:"maximumNumberOfPresetTours"`
	PTZPresetTourOperation     string            `bson:"ptzPresetTourOperation" json:"ptzPresetTourOperation"`
	Extension                  map[string]string `bson:"extension" json:"extension"`
}

type PTZConfigurationOptions struct {
	PTZRamps           IntAttrList               `bson:"ptzRamps" json:"ptzRamps"`
	Spaces             PTZSpaces                 `bson:"spaces" json:"spaces"`
	PTZTimeout         DurationRange             `bson:"ptzTimeout" json:"ptzTimeout"`
	PTControlDirection PTControlDirectionOptions `bson:"ptControlDirection" json:"ptControlDirection"`
	Extension          map[string]string         `bson:"extension" json:"extension"`
}

type IntAttrList struct {
	IntAttrList []int `bson:"intAttrList" json:"intAttrList"`
}

type DurationRange struct {
	Min TimeDuration `bson:"min" json:"min"`
	Max TimeDuration `bson:"max" json:"max"`
}

type PTControlDirectionOptions struct {
	EFlip     EFlipOptions      `bson:"eFlip" json:"eFlip"`
	Reverse   ReverseOptions    `bson:"reverse" json:"reverse"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type EFlipOptions struct {
	Mode      EFlipMode         `bson:"mode" json:"mode"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type ReverseOptions struct {
	Mode      string            `bson:"mode" json:"mode"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type PTZPreset struct {
	Token       onvif.ReferenceToken `bson:"token" json:"token"`
	Name        string               `bson:"name" json:"name"`
	PTZPosition PTZVector            `bson:"ptzPosition" json:"ptzPosition"`
}

type PTZVector struct {
	PanTilt Vector2D `bson:"panTilt" json:"panTilt"`
	Zoom    Vector1D `bson:"zoom" json:"zoom"`
}

type PTZStatus struct {
	Position   PTZVector     `bson:"position" json:"position"`
	MoveStatus PTZMoveStatus `bson:"moveStatus" json:"moveStatus"`
	Error      string        `bson:"error" json:"error"`
	UtcTime    time.Time     `bson:"utcTime" json:"utcTime"`
}
type PTZMoveStatus struct {
	PanTilt MoveStatus `bson:"panTilt" json:"panTilt"`
	Zoom    MoveStatus `bson:"zoom" json:"zoom"`
}

type MoveStatus struct {
	Status string `bson:"status" json:"status"`
}

type GeoLocation struct {
	Lon       float64 `bson:"lon" json:"lon"`
	Lat       float64 `bson:"lat" json:"lat"`
	Elevation float32 `bson:"elevation" json:"elevation"`
}

type PresetTour struct {
	Token             onvif.ReferenceToken           `bson:"token" json:"token"`
	Name              string                         `bson:"name" json:"name"`
	Status            PTZPresetTourStatus            `bson:"status" json:"status"`
	AutoStart         bool                           `bson:"autoStart" json:"autoStart"`
	StartingCondition PTZPresetTourStartingCondition `bson:"startingCondition" json:"startingCondition"`
	TourSpot          PTZPresetTourSpot              `bson:"tourSpot" json:"tourSpot"`
	Extension         map[string]string              `bson:"extension" json:"extension"`
}

type PTZPresetTourStatus struct {
	State           string            `bson:"state" json:"state"`
	CurrentTourSpot PTZPresetTourSpot `bson:"currentTourSpot" json:"currentTourSpot"`
	Extension       map[string]string `bson:"extension" json:"extension"`
}

type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail `bson:"presetDetail" json:"presetDetail"`
	Speed        PTZSpeed                  `bson:"speed" json:"speed"`
	StayTime     TimeDuration              `bson:"stayTime" json:"stayTime"`
	Extension    map[string]string         `bson:"extension" json:"extension"`
}

type PTZPresetTourPresetDetail struct {
	PresetToken   onvif.ReferenceToken `bson:"presetToken" json:"presetToken"`
	Home          bool                 `bson:"home" json:"home"`
	PTZPosition   PTZVector            `bson:"ptzPosition" json:"ptzPosition"`
	TypeExtension map[string]string    `bson:"typeExtension" json:"typeExtension"`
}

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder bool              `bson:"randomPresetOrder" json:"randomPresetOrder"`
	RecurringTime     int               `bson:"recurringTime" json:"recurringTime"`
	RecurringDuration TimeDuration      `bson:"recurringDuration" json:"recurringDuration"`
	Direction         string            `bson:"direction" json:"direction"`
	Extension         map[string]string `bson:"extension" json:"extension"`
}

type PTZPresetTourOptions struct {
	AutoStart         bool                                  `bson:"autoStart" json:"autoStart"`
	StartingCondition PTZPresetTourStartingConditionOptions `bson:"startingCondition" json:"startingCondition"`
	TourSpot          PTZPresetTourSpotOptions              `bson:"tourSpot" json:"tourSpot"`
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     IntRange          `bson:"recurringTime" json:"recurringTime"`
	RecurringDuration DurationRange     `bson:"recurringDuration" json:"recurringDuration"`
	Direction         string            `bson:"direction" json:"direction"`
	Extension         map[string]string `bson:"extension" json:"extension"`
}

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions `bson:"presetDetail" json:"presetDetail"`
	StayTime     DurationRange                    `bson:"stayTime" json:"stayTime"`
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          onvif.ReferenceToken `bson:"presetToken" json:"presetToken"`
	Home                 bool                 `bson:"home" json:"home"`
	PanTiltPositionSpace Space2DDescription   `bson:"panTiltPositionSpace" json:"panTiltPositionSpace"`
	ZoomPositionSpace    Space1DDescription   `bson:"zoomPositionSpace" json:"zoomPositionSpace"`
	Extension            map[string]string    `bson:"extension" json:"extension"`
}

// Capabilities of device
type Capabilities struct {
	Analytics AnalyticsCapabilities `bson:"analytics" json:"analytics"`
	Device    DeviceCapabilities    `bson:"device" json:"device"`
	Events    EventCapabilities     `bson:"events" json:"events"`
	Imaging   ImagingCapabilities   `bson:"imaging" json:"imaging"`
	Media     MediaCapabilities     `bson:"media" json:"media"`
	PTZ       PTZCapabilities       `bson:"ptz" json:"ptz"`
	Extension CapabilitiesExtension `bson:"extension" json:"extension"`
}

// AnalyticsCapabilities Check
type AnalyticsCapabilities struct {
	XAddr                  string `bson:"xAddr" json:"xAddr"`
	RuleSupport            bool   `bson:"ruleSupport" json:"ruleSupport"`
	AnalyticsModuleSupport bool   `bson:"analyticsModuleSupport" json:"analyticsModuleSupport"`
}

// DeviceCapabilities Check
type DeviceCapabilities struct {
	XAddr     string               `bson:"xAddr" json:"xAddr"`
	Network   NetworkCapabilities  `bson:"network" json:"network"`
	System    SystemCapabilities   `bson:"system" json:"system"`
	IO        IOCapabilities       `bson:"io" json:"io"`
	Security  SecurityCapabilities `bson:"security" json:"security"`
	Extension map[string]string    `bson:"extension" json:"extension"`
}

// NetworkCapabilities Check
type NetworkCapabilities struct {
	IPFilter          bool                         `bson:"ipFilter" json:"ipFilter"`
	ZeroConfiguration bool                         `bson:"zeroConfiguration" json:"zeroConfiguration"`
	IPVersion6        bool                         `bson:"ipVersion6" json:"ipVersion6"`
	DynDNS            bool                         `bson:"dynDNS" json:"dynDNS"`
	Extension         NetworkCapabilitiesExtension `bson:"extension" json:"extension"`
}

// NetworkCapabilitiesExtension Check
type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool              `bson:"dot11Configuration" json:"dot11Configuration"`
	Extension          map[string]string `bson:"extension" json:"extension"`
}

// SystemCapabilities check
type SystemCapabilities struct {
	DiscoveryResolve  bool              `bson:"discoveryResolve" json:"discoveryResolve"`
	DiscoveryBye      bool              `bson:"discoveryBye" json:"discoveryBye"`
	RemoteDiscovery   bool              `bson:"remoteDiscovery" json:"remoteDiscovery"`
	SystemBackup      bool              `bson:"systemBackup" json:"systemBackup"`
	SystemLogging     bool              `bson:"systemLogging" json:"systemLogging"`
	FirmwareUpgrade   bool              `bson:"firmwareUpgrade" json:"firmwareUpgrade"`
	SupportedVersions OnvifVersion      `bson:"supportedVersions" json:"supportedVersions"`
	Extension         map[string]string `bson:"extension" json:"extension"`
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    bool              `bson:"httpFirmwareUpgrade" json:"httpFirmwareUpgrade"`
	HttpSystemBackup       bool              `bson:"httpSystemBackup" json:"httpSystemBackup"`
	HttpSystemLogging      bool              `bson:"httpSystemLogging" json:"httpSystemLogging"`
	HttpSupportInformation bool              `bson:"httpSupportInformation" json:"httpSupportInformation"`
	Extension              map[string]string `bson:"extension" json:"extension"`
}

type IOCapabilities struct {
	InputConnectors int                     `bson:"inputConnectors" json:"inputConnectors"`
	RelayOutputs    int                     `bson:"relayOutputs" json:"relayOutputs"`
	Extension       IOCapabilitiesExtension `bson:"extension" json:"extension"`
}

type IOCapabilitiesExtension struct {
	Auxiliary         bool              `bson:"auxiliary" json:"auxiliary"`
	AuxiliaryCommands string            `bson:"auxiliaryCommands" json:"auxiliaryCommands"`
	Extension         map[string]string `bson:"extension" json:"extension"`
}

type SecurityCapabilities struct {
	TLS1_1               bool                          `bson:"tls1_1" json:"tls1_1"`
	TLS1_2               bool                          `bson:"tls1_2" json:"tls1_2"`
	OnboardKeyGeneration bool                          `bson:"onboardKeyGeneration" json:"onboardKeyGeneration"`
	AccessPolicyConfig   bool                          `bson:"accessPolicyConfig" json:"accessPolicyConfig"`
	X_509Token           bool                          `bson:"x_509Token" json:"x_509Token"`
	SAMLToken            bool                          `bson:"samlToken" json:"samlToken"`
	KerberosToken        bool                          `bson:"kerberosToken" json:"kerberosToken"`
	RELToken             bool                          `bson:"relToken" json:"relToken"`
	Extension            SecurityCapabilitiesExtension `bson:"extension" json:"extension"`
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    bool                           `bson:"tls1_0" json:"tls1_0"`
	Extension SecurityCapabilitiesExtension2 `bson:"extension" json:"extension"`
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              bool `bson:"dot1x" json:"dot1x"`
	SupportedEAPMethod int  `bson:"supportedEAPMethod" json:"supportedEAPMethod"`
	RemoteUserHandling bool `bson:"remoteUserHandling" json:"remoteUserHandling"`
}

type EventCapabilities struct {
	XAddr                                         string `bson:"xAddr" json:"xAddr"`
	WSSubscriptionPolicySupport                   bool   `bson:"wsSubscriptionPolicySupport" json:"wsSubscriptionPolicySupport"`
	WSPullPointSupport                            bool   `bson:"wsPullPointSupport" json:"wsPullPointSupport"`
	WSPausableSubscriptionManagerInterfaceSupport bool   `bson:"wsPausableSubscriptionManagerInterfaceSupport" json:"wsPausableSubscriptionManagerInterfaceSupport"`
}

type ImagingCapabilities struct {
	XAddr string `bson:"xAddr" json:"xAddr"`
}

type MediaCapabilities struct {
	XAddr                 string                        `bson:"xAddr" json:"xAddr"`
	StreamingCapabilities RealTimeStreamingCapabilities `bson:"streamingCapabilities" json:"streamingCapabilities"`
	Extension             MediaCapabilitiesExtension    `bson:"extension" json:"extension"`
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast bool              `bson:"rtpMulticast" json:"rtpMulticast"`
	RTP_TCP      bool              `bson:"rtp_tcp" json:"rtp_tcp"`
	RTP_RTSP_TCP bool              `bson:"rtp_rtsp_tcp" json:"rtp_rtsp_tcp"`
	Extension    map[string]string `bson:"extension" json:"extension"`
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities `bson:"profileCapabilities" json:"profileCapabilities"`
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `bson:"maximumNumberOfProfiles" json:"maximumNumberOfProfiles"`
}

type PTZCapabilities struct {
	XAddr string `bson:"xAddr" json:"xAddr"`
}

type CapabilitiesExtension struct {
	DeviceIO        DeviceIOCapabilities        `bson:"deviceIO" json:"deviceIO"`
	Display         DisplayCapabilities         `bson:"display" json:"display"`
	Recording       RecordingCapabilities       `bson:"recording" json:"recording"`
	Search          SearchCapabilities          `bson:"search" json:"search"`
	Replay          ReplayCapabilities          `bson:"replay" json:"replay"`
	Receiver        ReceiverCapabilities        `bson:"receiver" json:"receiver"`
	AnalyticsDevice AnalyticsDeviceCapabilities `bson:"analyticsDevice" json:"analyticsDevice"`
	Extensions      map[string]string           `bson:"extensions" json:"extensions"`
}
type DeviceIOCapabilities struct {
	XAddr        string `bson:"xAddr" json:"xAddr"`
	VideoSources int    `bson:"videoSources" json:"videoSources"`
	VideoOutputs int    `bson:"videoOutputs" json:"videoOutputs"`
	AudioSources int    `bson:"audioSources" json:"audioSources"`
	AudioOutputs int    `bson:"audioOutputs" json:"audioOutputs"`
	RelayOutputs int    `bson:"relayOutputs" json:"relayOutputs"`
}

type DisplayCapabilities struct {
	XAddr       string `bson:"xAddr" json:"xAddr"`
	FixedLayout bool   `bson:"fixedLayout" json:"fixedLayout"`
}

type RecordingCapabilities struct {
	XAddr              string `bson:"xAddr" json:"xAddr"`
	ReceiverSource     bool   `bson:"receiverSource" json:"receiverSource"`
	MediaProfileSource bool   `bson:"mediaProfileSource" json:"mediaProfileSource"`
	DynamicRecordings  bool   `bson:"dynamicRecordings" json:"dynamicRecordings"`
	DynamicTracks      bool   `bson:"dynamicTracks" json:"dynamicTracks"`
	MaxStringLength    int    `bson:"maxStringLength" json:"maxStringLength"`
}

type SearchCapabilities struct {
	XAddr          string `bson:"xAddr" json:"xAddr"`
	MetadataSearch bool   `bson:"metadataSearch" json:"metadataSearch"`
}

type ReplayCapabilities struct {
	XAddr string `bson:"xAddr" json:"xAddr"`
}

type ReceiverCapabilities struct {
	XAddr                string `bson:"xAddr" json:"xAddr"`
	RTP_Multicast        bool   `bson:"rtpMulticast" json:"rtpMulticast"`
	RTP_TCP              bool   `bson:"rtp_tcp" json:"rtp_tcp"`
	RTP_RTSP_TCP         bool   `bson:"rtp_rtsp_tcp" json:"rtp_rtsp_tcp"`
	SupportedReceivers   int    `bson:"supportedReceivers" json:"supportedReceivers"`
	MaximumRTSPURILength int    `bson:"maximumRTSPURILength" json:"maximumRTSPURILength"`
}

type AnalyticsDeviceCapabilities struct {
	XAddr       string            `bson:"xAddr" json:"xAddr"`
	RuleSupport bool              `bson:"ruleSupport" json:"ruleSupport"`
	Extension   map[string]string `bson:"extension" json:"extension"`
}

type HostnameInformation struct {
	FromDHCP  bool              `bson:"fromDHCP" json:"fromDHCP"`
	Name      string            `bson:"name" json:"name"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type DNSInformation struct {
	FromDHCP     bool              `bson:"fromDHCP" json:"fromDHCP"`
	SearchDomain string            `bson:"searchDomain" json:"searchDomain"`
	DNSFromDHCP  IPAddress         `bson:"dnsFromDHCP" json:"dnsFromDHCP"`
	DNSManual    IPAddress         `bson:"dnsManual" json:"dnsManual"`
	Extension    map[string]string `bson:"extension" json:"extension"`
}

type NTPInformation struct {
	FromDHCP    bool              `bson:"fromDHCP" json:"fromDHCP"`
	NTPFromDHCP string            `bson:"ntpFromDHCP" json:"ntpFromDHCP"`
	NTPManual   string            `bson:"ntpManual" json:"ntpManual"`
	Extension   map[string]string `bson:"extension" json:"extension"`
}

type DynamicDNSInformation struct {
	Type      string            `bson:"type" json:"type"`
	Name      string            `bson:"name" json:"name"`
	TTL       TimeDuration      `bson:"ttl" json:"ttl"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type NetworkInterface struct {
	DeviceEntity
	Enabled   bool                      `bson:"enabled" json:"enabled"`
	Info      NetworkInterfaceInfo      `bson:"info" json:"info"`
	Link      NetworkInterfaceLink      `bson:"link" json:"link"`
	IPv4      IPv4NetworkInterface      `bson:"ipv4" json:"ipv4"`
	IPv6      IPv6NetworkInterface      `bson:"ipv6" json:"ipv6"`
	Extension NetworkInterfaceExtension `bson:"extension" json:"extension"`
}

type NetworkInterfaceInfo struct {
	Name      string `bson:"name" json:"name"`
	HwAddress string `bson:"hwAddress" json:"hwAddress"`
	MTU       int    `bson:"mtu" json:"mtu"`
}

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting `bson:"adminSettings" json:"adminSettings"`
	OperSettings  NetworkInterfaceConnectionSetting `bson:"operSettings" json:"operSettings"`
	InterfaceType int                               `bson:"interfaceType" json:"interfaceType"`
}

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation bool   `bson:"autoNegotiation" json:"autoNegotiation"`
	Speed           int    `bson:"speed" json:"speed"`
	Duplex          string `bson:"duplex" json:"duplex"`
}

type NetworkInterfaceExtension struct {
	InterfaceType int                `bson:"interfaceType" json:"interfaceType"`
	Dot3          string             `bson:"dot3" json:"dot3"`
	Dot11         Dot11Configuration `bson:"dot11" json:"dot11"`
	Extension     map[string]string  `bson:"extension" json:"extension"`
}

type Dot11Configuration struct {
	SSID     string                     `bson:"ssid" json:"ssid"`
	Mode     string                     `bson:"mode" json:"mode"`
	Alias    string                     `bson:"alias" json:"alias"`
	Priority int64                      `bson:"priority" json:"priority"`
	Security Dot11SecurityConfiguration `bson:"security" json:"security"`
}

type Dot11SecurityConfiguration struct {
	Mode      string               `bson:"mode" json:"mode"`
	Algorithm string               `bson:"algorithm" json:"algorithm"`
	PSK       Dot11PSKSet          `bson:"psk" json:"psk"`
	Dot1X     onvif.ReferenceToken `bson:"dot1x" json:"dot1x"`
	Extension map[string]string    `bson:"extension" json:"extension"`
}

type Dot11PSKSet struct {
	Key        string            `bson:"key" json:"key"`
	Passphrase string            `bson:"passphrase" json:"passphrase"`
	Extension  map[string]string `bson:"extension" json:"extension"`
}

type IPv6NetworkInterface struct {
	Enabled bool              `bson:"enabled" json:"enabled"`
	Config  IPv6Configuration `bson:"config" json:"config"`
}

type IPv6Configuration struct {
	AcceptRouterAdvert bool                `bson:"acceptRouterAdvert" json:"acceptRouterAdvert"`
	DHCP               string              `bson:"dhcp" json:"dhcp"`
	Manual             PrefixedIPv6Address `bson:"manual" json:"manual"`
	LinkLocal          PrefixedIPv6Address `bson:"linkLocal" json:"linkLocal"`
	FromDHCP           PrefixedIPv6Address `bson:"fromDHCP" json:"fromDHCP"`
	FromRA             PrefixedIPv6Address `bson:"fromRA" json:"fromRA"`
	Extension          map[string]string   `bson:"extension" json:"extension"`
}

type PrefixedIPv6Address struct {
	Address      string `bson:"address" json:"address"`
	PrefixLength int    `bson:"prefixLength" json:"prefixLength"`
}

type IPv4NetworkInterface struct {
	Enabled bool              `bson:"enabled" json:"enabled"`
	Config  IPv4Configuration `bson:"config" json:"config"`
}

type IPv4Configuration struct {
	Manual    PrefixedIPv4Address `bson:"manual" json:"manual"`
	LinkLocal PrefixedIPv4Address `bson:"linkLocal" json:"linkLocal"`
	FromDHCP  PrefixedIPv4Address `bson:"fromDHCP" json:"fromDHCP"`
	DHCP      bool                `bson:"dhcp" json:"dhcp"`
}

type PrefixedIPv4Address struct {
	Address      string `bson:"address" json:"address"`
	PrefixLength int    `bson:"prefixLength" json:"prefixLength"`
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   bool                                      `bson:"enabled" json:"enabled"`
	Link      NetworkInterfaceConnectionSetting         `bson:"link" json:"link"`
	MTU       int                                       `bson:"mtu" json:"mtu"`
	IPv4      IPv4NetworkInterfaceSetConfiguration      `bson:"ipv4" json:"ipv4"`
	IPv6      IPv6NetworkInterfaceSetConfiguration      `bson:"ipv6" json:"ipv6"`
	Extension NetworkInterfaceSetConfigurationExtension `bson:"extension" json:"extension"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      string             `bson:"dot3" json:"dot3"`
	Dot11     Dot11Configuration `bson:"dot11" json:"dot11"`
	Extension map[string]string  `bson:"extension" json:"extension"`
}

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            bool                `bson:"enabled" json:"enabled"`
	AcceptRouterAdvert bool                `bson:"acceptRouterAdvert" json:"acceptRouterAdvert"`
	Manual             PrefixedIPv6Address `bson:"manual" json:"manual"`
	DHCP               string              `bson:"dhcp" json:"dhcp"`
}
type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled bool                `bson:"enabled" json:"enabled"`
	Manual  PrefixedIPv4Address `bson:"manual" json:"manual"`
	DHCP    bool                `bson:"dhcp" json:"dhcp"`
}

type NetworkProtocol struct {
	Name      string            `bson:"name" json:"name"`
	Enabled   bool              `bson:"enabled" json:"enabled"`
	Port      int               `bson:"port" json:"port"`
	Extension map[string]string `bson:"extension" json:"extension"`
}

type NetworkGateway struct {
	IPv4Address string `bson:"ipv4Address" json:"ipv4Address"`
	IPv6Address string `bson:"ipv6Address" json:"ipv6Address"`
}

type NetworkZeroConfiguration struct {
	InterfaceToken onvif.ReferenceToken              `bson:"interfaceToken" json:"interfaceToken"`
	Enabled        bool                              `bson:"enabled" json:"enabled"`
	Addresses      string                            `bson:"addresses" json:"addresses"`
	Extension      NetworkZeroConfigurationExtension `bson:"extension" json:"extension"`
}

type NetworkZeroConfigurationExtension struct {
	Additional *NetworkZeroConfiguration `bson:"additional" json:"additional"`
	Extension  map[string]string         `bson:"extension" json:"extension"`
}

type IPAddressFilter struct {
	Type        string              `bson:"type" json:"type"`
	IPv4Address PrefixedIPv4Address `bson:"ipv4Address" json:"ipv4Address"`
	IPv6Address PrefixedIPv6Address `bson:"ipv6Address" json:"ipv6Address"`
	Extension   map[string]string   `bson:"extension" json:"extension"`
}

type BinaryData struct {
	X    string `bson:"x" json:"x"`
	Data []byte `bson:"data" json:"data"`
}

type Certificate struct {
	CertificateID string     `bson:"certificateID" json:"certificateID"`
	Certificate   BinaryData `bson:"certificate" json:"certificate"`
}

type CertificateStatus struct {
	CertificateID string `bson:"certificateID" json:"certificateID"`
	Status        bool   `bson:"status" json:"status"`
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings `bson:"properties" json:"properties"`
}

type RelayOutputSettings struct {
	Mode      string       `bson:"mode" json:"mode"`
	DelayTime TimeDuration `bson:"delayTime" json:"delayTime"`
	IdleState string       `bson:"idleState" json:"idleState"`
}

type CertificateWithPrivateKey struct {
	CertificateID string     `bson:"certificateID" json:"certificateID"`
	Certificate   BinaryData `bson:"certificate" json:"certificate"`
	PrivateKey    BinaryData `bson:"privateKey" json:"privateKey"`
}

type CertificateInformation struct {
	CertificateID      string            `bson:"certificateID" json:"certificateID"`
	IssuerDN           string            `bson:"issuerDN" json:"issuerDN"`
	SubjectDN          string            `bson:"subjectDN" json:"subjectDN"`
	KeyUsage           CertificateUsage  `bson:"keyUsage" json:"keyUsage"`
	ExtendedKeyUsage   CertificateUsage  `bson:"extendedKeyUsage" json:"extendedKeyUsage"`
	KeyLength          int               `bson:"keyLength" json:"keyLength"`
	Version            string            `bson:"version" json:"version"`
	SerialNum          string            `bson:"serialNum" json:"serialNum"`
	SignatureAlgorithm string            `bson:"signatureAlgorithm" json:"signatureAlgorithm"`
	Validity           DateTimeRange     `bson:"validity" json:"validity"`
	Extension          map[string]string `bson:"extension" json:"extension"`
}

type DateTimeRange struct {
	From  time.Time `bson:"from" json:"from"`
	Until time.Time `bson:"until" json:"until"`
}

type CertificateUsage struct {
	Critical         bool   `bson:"critical" json:"critical"`
	CertificateUsage string `bson:"certificateUsage" json:"certificateUsage"`
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken onvif.ReferenceToken   `bson:"dot1xConfigurationToken" json:"dot1xConfigurationToken"`
	Identity                string                 `bson:"identity" json:"identity"`
	AnonymousID             string                 `bson:"anonymousID" json:"anonymousID"`
	EAPMethod               int                    `bson:"eapMethod" json:"eapMethod"`
	CACertificateID         string                 `bson:"caCertificateID" json:"caCertificateID"`
	EAPMethodConfiguration  EAPMethodConfiguration `bson:"eapMethodConfiguration" json:"eapMethodConfiguration"`
	Extension               map[string]string      `bson:"extension" json:"extension"`
}

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration  `bson:"tlsConfiguration" json:"tlsConfiguration"`
	Password         string            `bson:"password" json:"password"`
	Extension        map[string]string `bson:"extension" json:"extension"`
}

type TLSConfiguration struct {
	CertificateID string `bson:"certificateID" json:"certificateID"`
}

type Dot11Capabilities struct {
	TKIP                  bool `bson:"tkip" json:"tkip"`
	ScanAvailableNetworks bool `bson:"scanAvailableNetworks" json:"scanAvailableNetworks"`
	MultipleConfiguration bool `bson:"multipleConfiguration" json:"multipleConfiguration"`
	AdHocStationMode      bool `bson:"adHocStationMode" json:"adHocStationMode"`
	WEP                   bool `bson:"wep" json:"wep"`
}

type Dot11Status struct {
	SSID              string               `bson:"ssid" json:"ssid"`
	BSSID             string               `bson:"bssid" json:"bssid"`
	PairCipher        string               `bson:"pairCipher" json:"pairCipher"`
	GroupCipher       string               `bson:"groupCipher" json:"groupCipher"`
	SignalStrength    string               `bson:"signalStrength" json:"signalStrength"`
	ActiveConfigAlias onvif.ReferenceToken `bson:"activeConfigAlias" json:"activeConfigAlias"`
}
type Dot11AvailableNetworks struct {
	SSID                  string            `bson:"ssid" json:"ssid"`
	BSSID                 string            `bson:"bssid" json:"bssid"`
	AuthAndMangementSuite string            `bson:"authAndMangementSuite" json:"authAndMangementSuite"`
	PairCipher            string            `bson:"pairCipher" json:"pairCipher"`
	GroupCipher           string            `bson:"groupCipher" json:"groupCipher"`
	SignalStrength        string            `bson:"signalStrength" json:"signalStrength"`
	Extension             map[string]string `bson:"extension" json:"extension"`
}

type SystemLogUriList struct {
	SystemLog SystemLogUri `bson:"systemLog" json:"systemLog"`
}

type SystemLogUri struct {
	Type string `bson:"type" json:"type"`
	Uri  string `bson:"uri" json:"uri"`
}

type LocationEntity struct {
	Entity           string               `bson:"entity" json:"entity"`
	Token            onvif.ReferenceToken `bson:"token" json:"token"`
	Fixed            bool                 `bson:"fixed" json:"fixed"`
	GeoSource        string               `bson:"geoSource" json:"geoSource"`
	AutoGeo          bool                 `bson:"autoGeo" json:"autoGeo"`
	GeoLocation      GeoLocation          `bson:"geoLocation" json:"geoLocation"`
	GeoOrientation   GeoOrientation       `bson:"geoOrientation" json:"geoOrientation"`
	LocalLocation    LocalLocation        `bson:"localLocation" json:"localLocation"`
	LocalOrientation LocalOrientation     `bson:"localOrientation" json:"localOrientation"`
}

type LocalOrientation struct {
	Lon       float64 `bson:"lon" json:"lon"`
	Lat       float64 `bson:"lat" json:"lat"`
	Elevation float32 `bson:"elevation" json:"elevation"`
}

type LocalLocation struct {
	X float32 `bson:"x" json:"x"`
	Y float32 `bson:"y" json:"y"`
	Z float32 `bson:"z" json:"z"`
}

type GeoOrientation struct {
	Roll  float32 `bson:"roll" json:"roll"`
	Pitch float32 `bson:"pitch" json:"pitch"`
	Yaw   float32 `bson:"yaw" json:"yaw"`
}

type FocusMove struct {
	Absolute   AbsoluteFocus   `bson:"absolute" json:"absolute"`
	Relative   RelativeFocus   `bson:"relative" json:"relative"`
	Continuous ContinuousFocus `bson:"continuous" json:"continuous"`
}

type ContinuousFocus struct {
	Speed float32 `bson:"speed" json:"speed"`
}

type RelativeFocus struct {
	Distance float32 `bson:"distance" json:"distance"`
	Speed    float32 `bson:"speed" json:"speed"`
}

type AbsoluteFocus struct {
	Position float32 `bson:"position" json:"position"`
	Speed    float32 `bson:"speed" json:"speed"`
}

type DateTime struct {
	Time Time `bson:"time" json:"time"`
	Date Date `bson:"date" json:"date"`
}

type Time struct {
	Hour   int `bson:"hour" json:"hour"`
	Minute int `bson:"minute" json:"minute"`
	Second int `bson:"second" json:"second"`
}

type Date struct {
	Year  int `bson:"year" json:"year"`
	Month int `bson:"month" json:"month"`
	Day   int `bson:"day" json:"day"`
}
