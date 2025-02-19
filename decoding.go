package onvif

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/jjbarbosa7/onvif/device"
	"github.com/jjbarbosa7/onvif/xsd/onvif"
)

func getTimeFromXsdDateTime(xsdDateTime *etree.Element) (time.Time, error) {
	// Find the <Time> element
	timeElement := xsdDateTime.FindElement("./Time")
	if timeElement == nil {
		return time.Time{}, fmt.Errorf("time element not found in xds.DateTime")
	}

	hourStr := strings.TrimSpace(timeElement.FindElement("Hour").Text())
	minStr := strings.TrimSpace(timeElement.FindElement("Minute").Text())
	secStr := strings.TrimSpace(timeElement.FindElement("Second").Text())

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return time.Time{}, err
	}
	min, err := strconv.Atoi(minStr)
	if err != nil {
		return time.Time{}, err
	}
	sec, err := strconv.Atoi(secStr)
	if err != nil {
		return time.Time{}, err
	}

	// Find the <Date> element
	dateElement := xsdDateTime.FindElement("./Date")
	if dateElement == nil {
		return time.Time{}, fmt.Errorf("date element not found in UTCDateTime")
	}
	yearStr := strings.TrimSpace(dateElement.FindElement("Year").Text())
	monthStr := strings.TrimSpace(dateElement.FindElement("Month").Text())
	dayStr := strings.TrimSpace(dateElement.FindElement("Day").Text())

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return time.Time{}, err
	}

	utcDateTime := time.Date(year, time.Month(monthInt), day, hour, min, sec, 0, time.UTC)

	return utcDateTime, nil
}
func getTimeDurationFromXsdDuration(xsdDuration *etree.Element) (device.TimeDuration, error) {
	// Find the <Time> element
	timeElement := xsdDuration.FindElement("./Time")
	if timeElement == nil {
		return "", fmt.Errorf("time element not found in xds.Duration")
	}

	hourStr := strings.TrimSpace(timeElement.FindElement("Hour").Text())
	minStr := strings.TrimSpace(timeElement.FindElement("Minute").Text())
	secStr := strings.TrimSpace(timeElement.FindElement("Second").Text())

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return "", err
	}
	min, err := strconv.Atoi(minStr)
	if err != nil {
		return "", err
	}
	sec, err := strconv.Atoi(secStr)
	if err != nil {
		return "", err
	}

	duration := time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second

	timeDuration := device.TimeDuration(duration.String())

	return timeDuration, nil
}

func (dev *Device) DecodeSystemDateTime(data []byte) (*device.SystemDateTime, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	systemDateTime := device.SystemDateTime{}

	if e := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/DateTimeType"); e != nil {
		systemDateTime.DateTimeType = e.Text()
	}
	if e := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/DaylightSavings"); e != nil {
		systemDateTime.DaylightSavings = e.Text() == "true"
	}
	if e := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/TimeZone"); e != nil {
		systemDateTime.TimeZone = e.Text()
	}
	if e := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/UTCDateTime"); e != nil {
		systemDateTime.UTCDateTime, _ = getTimeFromXsdDateTime(e)
	}
	if e := doc.FindElement("./Envelope/Body/GetSystemDateAndTimeResponse/SystemDateAndTime/LocalDateTime"); e != nil {
		systemDateTime.LocalDateTime, _ = getTimeFromXsdDateTime(e)
	}

	return &systemDateTime, nil
}

func (dev *Device) DecodeCapabilities(data []byte) (*device.Capabilities, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	capabilities := device.Capabilities{}

	// Analytics
	analytics := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Analytics")
	if analytics != nil {
		if e := analytics.FindElement("XAddr"); e != nil {
			capabilities.Analytics.XAddr = e.Text()
		}
		if e := analytics.FindElement("RuleSupport"); e != nil {
			capabilities.Analytics.RuleSupport = e.Text() == "true"
		}
		if e := analytics.FindElement("AnalyticsModuleSupport"); e != nil {
			capabilities.Analytics.AnalyticsModuleSupport = e.Text() == "true"
		}
	}

	// Device
	device := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Device")
	if device != nil {
		if e := device.FindElement("XAddr"); e != nil {
			capabilities.Device.XAddr = e.Text()
		}
		if e := device.FindElement("Network"); e != nil {
			if e1 := e.FindElement("IPFilter"); e1 != nil {
				capabilities.Device.Network.IPFilter = e1.Text() == "true"
			}
			if e1 := e.FindElement("ZeroConfiguration"); e1 != nil {
				capabilities.Device.Network.ZeroConfiguration = e1.Text() == "true"
			}
			if e1 := e.FindElement("IPVersion6"); e1 != nil {
				capabilities.Device.Network.IPVersion6 = e1.Text() == "true"
			}
			if e1 := e.FindElement("DynDNS"); e1 != nil {
				capabilities.Device.Network.DynDNS = e1.Text() == "true"
			}
		}

		if e := device.FindElement("System"); e != nil {
			if e1 := e.FindElement("DiscoveryResolve"); e1 != nil {
				capabilities.Device.System.DiscoveryResolve = e1.Text() == "true"
			}
			if e1 := e.FindElement("DiscoveryBye"); e1 != nil {
				capabilities.Device.System.DiscoveryBye = e1.Text() == "true"
			}
			if e1 := e.FindElement("RemoteDiscovery"); e1 != nil {
				capabilities.Device.System.RemoteDiscovery = e1.Text() == "true"
			}
			if e1 := e.FindElement("SystemBackup"); e1 != nil {
				capabilities.Device.System.SystemBackup = e1.Text() == "true"
			}
			if e1 := e.FindElement("SystemLogging"); e1 != nil {
				capabilities.Device.System.SystemLogging = e1.Text() == "true"
			}
			if e1 := e.FindElement("FirmwareUpgrade"); e1 != nil {
				capabilities.Device.System.FirmwareUpgrade = e1.Text() == "true"
			}
			if e1 := e.FindElement("SupportedVersions"); e1 != nil {
				capabilities.Device.System.SupportedVersions.Major, _ = strconv.Atoi(e1.FindElement("Major").Text())
				capabilities.Device.System.SupportedVersions.Minor, _ = strconv.Atoi(e1.FindElement("Minor").Text())
			}
		}

		if e := device.FindElement("IO"); e != nil {
			if e1 := e.FindElement("InputConnectors"); e1 != nil {
				capabilities.Device.IO.InputConnectors, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("RelayOutputs"); e1 != nil {
				capabilities.Device.IO.RelayOutputs, _ = strconv.Atoi(e1.Text())
			}
		}

		if e := device.FindElement("Security"); e != nil {
			if e1 := e.FindElement("TLS1.1"); e1 != nil {
				capabilities.Device.Security.TLS1_1 = e1.Text() == "true"
			}
			if e1 := e.FindElement("TLS1.2"); e1 != nil {
				capabilities.Device.Security.TLS1_2 = e1.Text() == "true"
			}
			if e1 := e.FindElement("OnboardKeyGeneration"); e1 != nil {
				capabilities.Device.Security.OnboardKeyGeneration = e1.Text() == "true"
			}
			if e1 := e.FindElement("AccessPolicyConfig"); e1 != nil {
				capabilities.Device.Security.AccessPolicyConfig = e1.Text() == "true"
			}
			if e1 := e.FindElement("X_509Token"); e1 != nil {
				capabilities.Device.Security.X_509Token = e1.Text() == "true"
			}
			if e1 := e.FindElement("SAMLToken"); e1 != nil {
				capabilities.Device.Security.SAMLToken = e1.Text() == "true"
			}
			if e1 := e.FindElement("KerberosToken"); e1 != nil {
				capabilities.Device.Security.KerberosToken = e1.Text() == "true"
			}
			if e1 := e.FindElement("RELToken"); e1 != nil {
				capabilities.Device.Security.RELToken = e1.Text() == "true"
			}
		}
	}

	// Events
	events := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Events")
	if events != nil {
		if e := events.FindElement("XAddr"); e != nil {
			capabilities.Events.XAddr = e.Text()
		}
		if e := events.FindElement("WSSubscriptionPolicySupport"); e != nil {
			capabilities.Events.WSSubscriptionPolicySupport = e.Text() == "true"
		}
		if e := events.FindElement("WSPullPointSupport"); e != nil {
			capabilities.Events.WSPullPointSupport = e.Text() == "true"
		}
		if e := events.FindElement("WSPausableSubscriptionManagerInterfaceSupport"); e != nil {
			capabilities.Events.WSPausableSubscriptionManagerInterfaceSupport = e.Text() == "true"
		}
	}

	// Imaging
	imaging := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Imaging")
	if imaging != nil {
		if e := imaging.FindElement("XAddr"); e != nil {
			capabilities.Imaging.XAddr = e.Text()
		}
	}

	// Media
	media := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Media")
	if media != nil {
		if e := media.FindElement("XAddr"); e != nil {
			capabilities.Media.XAddr = e.Text()
		}
		if e := media.FindElement("StreamingCapabilities"); e != nil {
			if e1 := e.FindElement("RTPMulticast"); e1 != nil {
				capabilities.Media.StreamingCapabilities.RTPMulticast = e1.Text() == "true"
			}
			if e1 := e.FindElement("RTP_TCP"); e1 != nil {
				capabilities.Media.StreamingCapabilities.RTP_TCP = e1.Text() == "true"
			}
			if e1 := e.FindElement("RTP_RTSP_TCP"); e1 != nil {
				capabilities.Media.StreamingCapabilities.RTP_RTSP_TCP = e1.Text() == "true"
			}
		}
	}

	// PTZ
	ptz := doc.FindElement("./Envelope/Body/GetCapabilitiesResponse/Capabilities/PTZ")
	if ptz != nil {
		if e := ptz.FindElement("XAddr"); e != nil {
			capabilities.PTZ.XAddr = e.Text()
		}
	}

	return &capabilities, nil
}

func (dev *Device) DecodePTZNode(data []byte) (*device.PTZNode, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	ptzNode := device.PTZNode{}

	node := doc.FindElement("./Envelope/Body/GetNodesResponse/PTZNode")
	if node == nil {
		return nil, fmt.Errorf("PTZNode element not found")
	}

	token := node.SelectAttrValue("token", "")
	tokenString := fmt.Sprintf("%v", token)
	ptzNode.Token = onvif.ReferenceToken(tokenString)
	if e := node.FindElement("Name"); e != nil {
		ptzNode.Name = e.Text()
	}
	if e := node.FindElement("FixedHomePosition"); e != nil {
		ptzNode.FixedHomePosition = e.Text() == "true"
	}
	if e := node.FindElement("node"); e != nil {
		ptzNode.GeoMove = e.Text() == "true"
	}
	if e := node.FindElement("SupportedPTZSpaces"); e != nil {
		if e1 := e.FindElement("AbsolutePanTiltPositionSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.AbsolutePanTiltPositionSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsolutePanTiltPositionSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsolutePanTiltPositionSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
			if e2 := e1.FindElement("YRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsolutePanTiltPositionSpace.YRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsolutePanTiltPositionSpace.YRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("AbsoluteZoomPositionSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.AbsoluteZoomPositionSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsoluteZoomPositionSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.AbsoluteZoomPositionSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("RelativePanTiltTranslationSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.RelativePanTiltTranslationSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativePanTiltTranslationSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativePanTiltTranslationSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
			if e2 := e1.FindElement("YRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativePanTiltTranslationSpace.YRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativePanTiltTranslationSpace.YRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("RelativeZoomTranslationSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.RelativeZoomTranslationSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativeZoomTranslationSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.RelativeZoomTranslationSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("ContinuousPanTiltVelocitySpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.ContinuousPanTiltVelocitySpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousPanTiltVelocitySpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousPanTiltVelocitySpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
			if e2 := e1.FindElement("YRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousPanTiltVelocitySpace.YRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousPanTiltVelocitySpace.YRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("ContinuousZoomVelocitySpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.ContinuousZoomVelocitySpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousZoomVelocitySpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ContinuousZoomVelocitySpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
		if e1 := e.FindElement("PanTiltSpeedSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.PanTiltSpeedSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.PanTiltSpeedSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.PanTiltSpeedSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}

		}
		if e1 := e.FindElement("ZoomSpeedSpace"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzNode.SupportedPTZSpaces.ZoomSpeedSpace.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ZoomSpeedSpace.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzNode.SupportedPTZSpaces.ZoomSpeedSpace.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
	}
	if e := node.FindElement("MaximumNumberOfPresets"); e != nil {
		ptzNode.MaximumNumberOfPresets, _ = strconv.Atoi(e.Text())
	}
	if e := node.FindElement("HomeSupported"); e != nil {
		ptzNode.HomeSupported = e.Text() == "true"
	}

	return &ptzNode, nil
}

func (dev *Device) DecodePTZConfiguration(data []byte) (*device.PTZConfiguration, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	ptzConfiguration := device.PTZConfiguration{}

	configuration := doc.FindElement("./Envelope/Body/GetConfigurationsResponse/PTZConfiguration")
	if configuration == nil {
		return nil, fmt.Errorf("PTZConfiguration element not found")
	}

	token := configuration.SelectAttrValue("token", "")
	tokenString := fmt.Sprintf("%v", token)
	ptzConfiguration.Token = onvif.ReferenceToken(tokenString)

	if e := configuration.FindElement("Name"); e != nil {
		ptzConfiguration.Name = e.Text()
	}
	if e := configuration.FindElement("UseCount"); e != nil {
		ptzConfiguration.UseCount, _ = strconv.Atoi(e.Text())
	}
	if e := configuration.FindElement("MoveRamp"); e != nil {
		ptzConfiguration.MoveRamp, _ = strconv.Atoi(e.Text())
	}
	if e := configuration.FindElement("PresetRamp"); e != nil {
		ptzConfiguration.PresetRamp, _ = strconv.Atoi(e.Text())
	}
	if e := configuration.FindElement("PresetTourRamp"); e != nil {
		ptzConfiguration.PresetTourRamp, _ = strconv.Atoi(e.Text())
	}
	if e := configuration.FindElement("NodeToken"); e != nil {
		ptzConfiguration.NodeToken = onvif.ReferenceToken(e.Text())
	}
	if e := configuration.FindElement("DefaultAbsolutePantTiltPositionSpace"); e != nil {
		ptzConfiguration.DefaultAbsolutePantTiltPositionSpace = e.Text()
	}
	if e := configuration.FindElement("DefaultAbsoluteZoomPositionSpace"); e != nil {
		ptzConfiguration.DefaultAbsoluteZoomPositionSpace = e.Text()
	}
	if e := configuration.FindElement("DefaultRelativePanTiltTranslationSpace"); e != nil {
		ptzConfiguration.DefaultRelativePanTiltTranslationSpace = e.Text()
	}
	if e := configuration.FindElement("DefaultRelativeZoomTranslationSpace"); e != nil {
		ptzConfiguration.DefaultRelativeZoomTranslationSpace = e.Text()
	}
	if e := configuration.FindElement("DefaultContinuousPanTiltVelocitySpace"); e != nil {
		ptzConfiguration.DefaultContinuousPanTiltVelocitySpace = e.Text()
	}
	if e := configuration.FindElement("DefaultContinuousZoomVelocitySpace"); e != nil {
		ptzConfiguration.DefaultContinuousZoomVelocitySpace = e.Text()
	}
	if e := configuration.FindElement("DefaultPTZSpeed"); e != nil {
		if e1 := e.FindElement("PanTilt"); e1 != nil {
			if e2 := e1.FindElement("x"); e2 != nil {
				ptzConfiguration.DefaultPTZSpeed.PanTilt.X, _ = strconv.ParseFloat(e2.Text(), 64)
			}
			if e2 := e1.FindElement("y"); e2 != nil {
				ptzConfiguration.DefaultPTZSpeed.PanTilt.Y, _ = strconv.ParseFloat(e2.Text(), 64)
			}
		}
		if e1 := e.FindElement("Zoom"); e1 != nil {
			if e2 := e1.FindElement("x"); e2 != nil {
				ptzConfiguration.DefaultPTZSpeed.Zoom.X, _ = strconv.ParseFloat(e2.Text(), 64)
			}
		}
	}
	if e := configuration.FindElement("DefaultPTZTimeout"); e != nil {
		ptzConfiguration.DefaultPTZTimeout, _ = getTimeDurationFromXsdDuration(e)
	}
	if e := configuration.FindElement("PanTiltLimits"); e != nil {
		if e1 := e.FindElement("Range"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzConfiguration.PanTiltLimits.Range.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzConfiguration.PanTiltLimits.Range.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzConfiguration.PanTiltLimits.Range.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
			if e2 := e1.FindElement("YRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzConfiguration.PanTiltLimits.Range.YRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzConfiguration.PanTiltLimits.Range.YRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
	}
	if e := configuration.FindElement("ZoomLimits"); e != nil {
		if e1 := e.FindElement("Range"); e1 != nil {
			if e2 := e1.FindElement("URI"); e2 != nil {
				ptzConfiguration.ZoomLimits.Range.URI = e2.Text()
			}
			if e2 := e1.FindElement("XRange"); e2 != nil {
				if e3 := e2.FindElement("Min"); e3 != nil {
					ptzConfiguration.ZoomLimits.Range.XRange.Min, _ = strconv.ParseFloat(e3.Text(), 64)
				}
				if e3 := e2.FindElement("Max"); e3 != nil {
					ptzConfiguration.ZoomLimits.Range.XRange.Max, _ = strconv.ParseFloat(e3.Text(), 64)
				}
			}
		}
	}

	return &ptzConfiguration, nil
}

func (dev *Device) DecodeProfiles(data []byte) ([]device.Profile, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	profiles := make([]device.Profile, 0)

	for _, profileElement := range doc.FindElements("./Envelope/Body/GetProfilesResponse/Profiles") {
		profile := device.Profile{}

		token := profileElement.SelectAttrValue("token", "")
		tokenString := fmt.Sprintf("%v", token)
		profile.Token = onvif.ReferenceToken(tokenString)

		fixed := profileElement.SelectAttrValue("fixed", "")
		fixedString := fmt.Sprintf("%v", fixed)
		profile.Fixed = fixedString == "true"

		if e := profileElement.FindElement("VideoSourceConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.VideoSourceConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.VideoSourceConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.VideoSourceConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("ViewMode"); e1 != nil {
				profile.VideoSourceConfiguration.ViewMode = e1.Text()
			}
			if e1 := e.FindElement("SourceToken"); e1 != nil {
				profile.VideoSourceConfiguration.SourceToken = onvif.ReferenceToken(e1.Text())
			}
			if e1 := e.FindElement("Bounds"); e1 != nil {
				if e2 := e1.FindElement("x"); e2 != nil {
					profile.VideoSourceConfiguration.Bounds.X, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("y"); e2 != nil {
					profile.VideoSourceConfiguration.Bounds.Y, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("width"); e2 != nil {
					profile.VideoSourceConfiguration.Bounds.Width, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("height"); e2 != nil {
					profile.VideoSourceConfiguration.Bounds.Height, _ = strconv.Atoi(e2.Text())
				}
			}
		}
		if e := profileElement.FindElement("AudioSourceConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.AudioSourceConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.AudioSourceConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.AudioSourceConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("SourceToken"); e1 != nil {
				profile.AudioSourceConfiguration.SourceToken = onvif.ReferenceToken(e1.Text())
			}
		}
		if e := profileElement.FindElement("VideoEncoderConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.VideoEncoderConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.VideoEncoderConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.VideoEncoderConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("Encoding"); e1 != nil {
				profile.VideoEncoderConfiguration.Encoding = e1.Text()
			}
			if e1 := e.FindElement("Resolution"); e1 != nil {
				if e2 := e1.FindElement("Width"); e2 != nil {
					profile.VideoEncoderConfiguration.Resolution.Width, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("Height"); e2 != nil {
					profile.VideoEncoderConfiguration.Resolution.Height, _ = strconv.Atoi(e2.Text())
				}
			}
			if e1 := e.FindElement("Quality"); e1 != nil {
				profile.VideoEncoderConfiguration.Quality, _ = strconv.ParseFloat(e1.Text(), 64)
			}
			if e1 := e.FindElement("RateControl"); e1 != nil {
				if e2 := e1.FindElement("FrameRateLimit"); e2 != nil {
					profile.VideoEncoderConfiguration.RateControl.FrameRateLimit, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("EncodingInterval"); e2 != nil {
					profile.VideoEncoderConfiguration.RateControl.EncodingInterval, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("BitrateLimit"); e2 != nil {
					profile.VideoEncoderConfiguration.RateControl.BitrateLimit, _ = strconv.Atoi(e2.Text())
				}
			}
			if e1 := e.FindElement("MPEG4"); e1 != nil {
				if e2 := e1.FindElement("GovLength"); e2 != nil {
					profile.VideoEncoderConfiguration.MPEG4.GovLength, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("Mpeg4Profile"); e2 != nil {
					profile.VideoEncoderConfiguration.MPEG4.Mpeg4Profile = e2.Text()
				}
			}
			if e1 := e.FindElement("H264"); e1 != nil {
				if e2 := e1.FindElement("GovLength"); e2 != nil {
					profile.VideoEncoderConfiguration.H264.GovLength, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("H264Profile"); e2 != nil {
					profile.VideoEncoderConfiguration.H264.H264Profile = e2.Text()
				}
			}
			if e1 := e.FindElement("Multicast"); e1 != nil {
				if e2 := e1.FindElement("Address"); e2 != nil {
					if e3 := e2.FindElement("Type"); e3 != nil {
						profile.VideoEncoderConfiguration.Multicast.Address.Type = e3.Text()
					}
					if e3 := e2.FindElement("IPv4Address"); e3 != nil {
						profile.VideoEncoderConfiguration.Multicast.Address.IPv4Address = e3.Text()
					}
					if e3 := e2.FindElement("IPv6Address"); e3 != nil {
						profile.VideoEncoderConfiguration.Multicast.Address.IPv6Address = e3.Text()
					}
				}
				if e2 := e1.FindElement("Port"); e2 != nil {
					profile.VideoEncoderConfiguration.Multicast.Port, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("TTL"); e2 != nil {
					profile.VideoEncoderConfiguration.Multicast.TTL, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("AutoStart"); e2 != nil {
					profile.VideoEncoderConfiguration.Multicast.AutoStart = e2.Text() == "true"
				}
			}
			if e1 := e.FindElement("SessionTimeout"); e1 != nil {
				sessionTimeout, err := getTimeDurationFromXsdDuration(e1)
				if err == nil {
					profile.VideoEncoderConfiguration.SessionTimeout = sessionTimeout
				}
			}
		}
		if e := profileElement.FindElement("AudioEncoderConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.AudioEncoderConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.AudioEncoderConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.AudioEncoderConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("Encoding"); e1 != nil {
				profile.AudioEncoderConfiguration.Encoding = e1.Text()
			}
			if e1 := e.FindElement("Bitrate"); e1 != nil {
				profile.AudioEncoderConfiguration.Bitrate, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("SampleRate"); e1 != nil {
				profile.AudioEncoderConfiguration.SampleRate, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("Multicast"); e1 != nil {
				if e2 := e1.FindElement("Address"); e2 != nil {
					if e3 := e2.FindElement("Type"); e3 != nil {
						profile.AudioEncoderConfiguration.Multicast.Address.Type = e3.Text()
					}
					if e3 := e2.FindElement("IPv4Address"); e3 != nil {
						profile.AudioEncoderConfiguration.Multicast.Address.IPv4Address = e3.Text()
					}
					if e3 := e2.FindElement("IPv6Address"); e3 != nil {
						profile.AudioEncoderConfiguration.Multicast.Address.IPv6Address = e3.Text()
					}
				}
				if e2 := e1.FindElement("Port"); e2 != nil {
					profile.AudioEncoderConfiguration.Multicast.Port, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("TTL"); e2 != nil {
					profile.AudioEncoderConfiguration.Multicast.TTL, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("AutoStart"); e2 != nil {
					profile.AudioEncoderConfiguration.Multicast.AutoStart = e2.Text() == "true"
				}
			}
			if e1 := e.FindElement("SessionTimeout"); e1 != nil {
				sessionTimeout, err := getTimeDurationFromXsdDuration(e1)
				if err == nil {
					profile.AudioEncoderConfiguration.SessionTimeout = sessionTimeout
				}
			}
		}
		if e := profileElement.FindElement("VideoAnalyticsConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.VideoAnalyticsConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.VideoAnalyticsConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.VideoAnalyticsConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("AnalyticsEngineConfiguration"); e1 != nil {
				if e2 := e1.FindElement("AnalyticsModule"); e2 != nil {
					if e3 := e2.FindElement("Name"); e3 != nil {
						profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Name = e3.Text()
					}
					if e3 := e2.FindElement("Type"); e3 != nil {
						if e4 := e3.FindElement("Namespace"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Type.Namespace = e4.Text()
						}
						if e4 := e3.FindElement("LocalPart"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Type.LocalPart = e4.Text()
						}
					}
					if e3 := e2.FindElement("Parameters"); e3 != nil {
						if e4 := e3.FindElement("SimpleItem"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.SimpleItem.Name = e4.SelectAttrValue("Name", "")
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.SimpleItem.Value = e4.Text()
						}
						if e4 := e3.FindElement("ElementItem"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.ElementItem.Name = e4.SelectAttrValue("Name", "")
						}
					}
				}

			}
			if e1 := e.FindElement("RuleEngineConfiguration"); e1 != nil {
				if e2 := e1.FindElement("Rule"); e2 != nil {
					if e3 := e2.FindElement("Name"); e3 != nil {
						profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Name = e3.Text()
					}
					if e3 := e2.FindElement("Type"); e3 != nil {
						if e4 := e3.FindElement("Namespace"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Type.Namespace = e4.Text()
						}
						if e4 := e3.FindElement("LocalPart"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Type.LocalPart = e4.Text()
						}
					}
					if e3 := e2.FindElement("Parameters"); e3 != nil {
						if e4 := e3.FindElement("SimpleItem"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.SimpleItem.Name = e4.SelectAttrValue("Name", "")
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.SimpleItem.Value = e4.Text()
						}
						if e4 := e3.FindElement("ElementItem"); e4 != nil {
							profile.VideoAnalyticsConfiguration.RuleEngineConfiguration.Rule.Parameters.ElementItem.Name = e4.SelectAttrValue("Name", "")
						}
					}
				}

			}
		}
		if e := profileElement.FindElement("PTZConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.PTZConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.PTZConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.PTZConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("NodeToken"); e1 != nil {
				profile.PTZConfiguration.NodeToken = onvif.ReferenceToken(e1.Text())
			}
			if e1 := e.FindElement("DefaultAbsolutePantTiltPositionSpace"); e1 != nil {
				profile.PTZConfiguration.DefaultAbsolutePantTiltPositionSpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultAbsoluteZoomPositionSpace"); e1 != nil {
				profile.PTZConfiguration.DefaultAbsoluteZoomPositionSpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultRelativePanTiltTranslationSpace"); e1 != nil {
				profile.PTZConfiguration.DefaultRelativePanTiltTranslationSpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultRelativeZoomTranslationSpace"); e1 != nil {
				profile.PTZConfiguration.DefaultRelativeZoomTranslationSpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultContinuousPanTiltVelocitySpace"); e1 != nil {
				profile.PTZConfiguration.DefaultContinuousPanTiltVelocitySpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultContinuousZoomVelocitySpace"); e1 != nil {
				profile.PTZConfiguration.DefaultContinuousZoomVelocitySpace = e1.Text()
			}
			if e1 := e.FindElement("DefaultPTZSpeed"); e1 != nil {
				if e2 := e1.FindElement("PanTilt"); e2 != nil {
					if e3 := e2.FindElement("space"); e3 != nil {
						profile.PTZConfiguration.DefaultPTZSpeed.PanTilt.Space = e3.Text()
					}
					if e3 := e2.FindElement("x"); e3 != nil {
						profile.PTZConfiguration.DefaultPTZSpeed.PanTilt.X, _ = strconv.ParseFloat(e3.Text(), 64)
					}
					if e3 := e2.FindElement("y"); e3 != nil {
						profile.PTZConfiguration.DefaultPTZSpeed.PanTilt.Y, _ = strconv.ParseFloat(e3.Text(), 64)
					}
				}
				if e2 := e1.FindElement("Zoom"); e2 != nil {
					if e3 := e2.FindElement("space"); e3 != nil {
						profile.PTZConfiguration.DefaultPTZSpeed.Zoom.Space = e3.Text()
					}
					if e3 := e2.FindElement("x"); e3 != nil {
						profile.PTZConfiguration.DefaultPTZSpeed.Zoom.X, _ = strconv.ParseFloat(e3.Text(), 64)
					}
				}
			}
			if e1 := e.FindElement("DefaultPTZTimeout"); e1 != nil {
				ptzTimeout, err := getTimeDurationFromXsdDuration(e1)
				if err == nil {
					profile.PTZConfiguration.DefaultPTZTimeout = ptzTimeout
				}
			}
			if e1 := e.FindElement("PanTiltLimits"); e1 != nil {
				if e2 := e1.FindElement("Range"); e2 != nil {
					if e3 := e2.FindElement("URI"); e3 != nil {
						profile.PTZConfiguration.PanTiltLimits.Range.URI = e3.Text()
					}
					if e3 := e2.FindElement("XRange"); e3 != nil {
						if e4 := e3.FindElement("Min"); e4 != nil {
							profile.PTZConfiguration.PanTiltLimits.Range.XRange.Min, _ = strconv.ParseFloat(e4.Text(), 64)
						}
						if e4 := e3.FindElement("Max"); e4 != nil {
							profile.PTZConfiguration.PanTiltLimits.Range.XRange.Max, _ = strconv.ParseFloat(e4.Text(), 64)
						}
					}
					if e3 := e2.FindElement("YRange"); e3 != nil {
						if e4 := e3.FindElement("Min"); e4 != nil {
							profile.PTZConfiguration.PanTiltLimits.Range.YRange.Min, _ = strconv.ParseFloat(e4.Text(), 64)
						}
						if e4 := e3.FindElement("Max"); e4 != nil {
							profile.PTZConfiguration.PanTiltLimits.Range.YRange.Max, _ = strconv.ParseFloat(e4.Text(), 64)
						}
					}
				}
			}
			if e1 := e.FindElement("ZoomLimits"); e1 != nil {
				if e2 := e1.FindElement("Range"); e2 != nil {
					if e3 := e2.FindElement("URI"); e3 != nil {
						profile.PTZConfiguration.ZoomLimits.Range.URI = e3.Text()
					}
					if e3 := e2.FindElement("XRange"); e3 != nil {
						if e4 := e3.FindElement("Min"); e4 != nil {
							profile.PTZConfiguration.ZoomLimits.Range.XRange.Min, _ = strconv.ParseFloat(e4.Text(), 64)
						}
						if e4 := e3.FindElement("Max"); e4 != nil {
							profile.PTZConfiguration.ZoomLimits.Range.XRange.Max, _ = strconv.ParseFloat(e4.Text(), 64)
						}
					}
				}
			}
		}
		if e := profileElement.FindElement("MetadataConfiguration"); e != nil {
			token := profileElement.SelectAttrValue("token", "")
			tokenString := fmt.Sprintf("%v", token)
			profile.MetadataConfiguration.Token = onvif.ReferenceToken(tokenString)

			if e1 := e.FindElement("Name"); e1 != nil {
				profile.MetadataConfiguration.Name = e1.Text()
			}
			if e1 := e.FindElement("UseCount"); e1 != nil {
				profile.MetadataConfiguration.UseCount, _ = strconv.Atoi(e1.Text())
			}
			if e1 := e.FindElement("CompressionType"); e1 != nil {
				profile.MetadataConfiguration.CompressionType = e1.Text()
			}
			if e1 := e.FindElement("PTZStatus"); e1 != nil {
				if e2 := e1.FindElement("Status"); e2 != nil {
					profile.MetadataConfiguration.PTZStatus.Status = e2.Text() == "true"
				}
				if e2 := e1.FindElement("Position"); e2 != nil {
					profile.MetadataConfiguration.PTZStatus.Position = e2.Text() == "true"
				}
			}
			if e1 := e.FindElement("Analytics"); e1 != nil {
				profile.MetadataConfiguration.Analytics = e1.Text() == "true"
			}
			if e1 := e.FindElement("Multicast"); e1 != nil {
				if e2 := e1.FindElement("Address"); e2 != nil {
					if e3 := e2.FindElement("Type"); e3 != nil {
						profile.MetadataConfiguration.Multicast.Address.Type = e3.Text()
					}
					if e3 := e2.FindElement("IPv4Address"); e3 != nil {
						profile.MetadataConfiguration.Multicast.Address.IPv4Address = e3.Text()
					}
					if e3 := e2.FindElement("IPv6Address"); e3 != nil {
						profile.MetadataConfiguration.Multicast.Address.IPv6Address = e3.Text()
					}
				}
				if e2 := e1.FindElement("Port"); e2 != nil {
					profile.MetadataConfiguration.Multicast.Port, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("TTL"); e2 != nil {
					profile.MetadataConfiguration.Multicast.TTL, _ = strconv.Atoi(e2.Text())
				}
				if e2 := e1.FindElement("AutoStart"); e2 != nil {
					profile.MetadataConfiguration.Multicast.AutoStart = e2.Text() == "true"
				}
			}
			if e1 := e.FindElement("SessionTimeout"); e1 != nil {
				sessionTimeout, err := getTimeDurationFromXsdDuration(e1)
				if err == nil {
					profile.MetadataConfiguration.SessionTimeout = sessionTimeout
				}
			}
			if e1 := e.FindElement("AnalyticsEngineConfiguration"); e1 != nil {
				if e2 := e1.FindElement("AnalyticsModule"); e2 != nil {
					if e3 := e2.FindElement("Name"); e3 != nil {
						profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Name = e3.Text()
					}
					if e3 := e2.FindElement("Type"); e3 != nil {
						if e4 := e3.FindElement("Namespace"); e4 != nil {
							profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Type.Namespace = e4.Text()
						}
						if e4 := e3.FindElement("LocalPart"); e4 != nil {
							profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Type.LocalPart = e4.Text()
						}
					}
					if e3 := e2.FindElement("Parameters"); e3 != nil {
						if e4 := e3.FindElement("SimpleItem"); e4 != nil {
							profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Parameters.SimpleItem.Name = e4.SelectAttrValue("Name", "")
							profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Parameters.SimpleItem.Value = e4.Text()
						}
						if e4 := e3.FindElement("ElementItem"); e4 != nil {
							profile.MetadataConfiguration.AnalyticsEngineConfiguration.AnalyticsModule.Parameters.ElementItem.Name = e4.SelectAttrValue("Name", "")
						}
					}
				}
			}
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (dev *Device) DecodePresets(data []byte) ([]device.PTZPreset, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	presets := make([]device.PTZPreset, 0)

	for _, presetElement := range doc.FindElements("./Envelope/Body/GetPresetsResponse/Preset") {
		preset := device.PTZPreset{}

		token := presetElement.SelectAttrValue("token", "")
		tokenString := fmt.Sprintf("%v", token)
		preset.Token = onvif.ReferenceToken(tokenString)

		if e := presetElement.FindElement("Name"); e != nil {
			preset.Name = e.Text()
		}
		if e := presetElement.FindElement("PTZPosition"); e != nil {
			if e1 := e.FindElement("PanTilt"); e1 != nil {
				space := e1.SelectAttrValue("space", "")
				x := e1.SelectAttrValue("x", "")
				y := e1.SelectAttrValue("y", "")

				preset.PTZPosition.PanTilt.Space = space
				preset.PTZPosition.PanTilt.X, _ = strconv.ParseFloat(x, 64)
				preset.PTZPosition.PanTilt.Y, _ = strconv.ParseFloat(y, 64)
			}
			if e1 := e.FindElement("Zoom"); e1 != nil {
				space := e1.SelectAttrValue("space", "")
				x := e1.SelectAttrValue("x", "")

				preset.PTZPosition.Zoom.Space = space
				preset.PTZPosition.Zoom.X, _ = strconv.ParseFloat(x, 64)
			}
		}

		presets = append(presets, preset)
	}

	return presets, nil
}

func (dev *Device) DecodeStatus(data []byte) (*device.PTZStatus, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	ptzStatus := device.PTZStatus{}

	status := doc.FindElement("./Envelope/Body/GetStatusResponse/PTZStatus")
	if status == nil {
		return nil, fmt.Errorf("PTZStatus element not found")
	}

	if e := status.FindElement("Position"); e != nil {
		if e1 := e.FindElement("PanTilt"); e1 != nil {
			space := e1.SelectAttrValue("space", "")
			x := e1.SelectAttrValue("x", "")
			y := e1.SelectAttrValue("y", "")

			ptzStatus.Position.PanTilt.Space = space
			ptzStatus.Position.PanTilt.X, _ = strconv.ParseFloat(x, 64)
			ptzStatus.Position.PanTilt.Y, _ = strconv.ParseFloat(y, 64)
		}
		if e1 := e.FindElement("Zoom"); e1 != nil {
			space := e1.SelectAttrValue("space", "")
			x := e1.SelectAttrValue("x", "")

			ptzStatus.Position.Zoom.Space = space
			ptzStatus.Position.Zoom.X, _ = strconv.ParseFloat(x, 64)
		}
	}
	if e := status.FindElement("MoveStatus"); e != nil {
		if e1 := e.FindElement("PanTilt"); e1 != nil {
			if e2 := e1.FindElement("Status"); e2 != nil {
				ptzStatus.MoveStatus.PanTilt.Status = e2.Text()
			} else {
				ptzStatus.MoveStatus.PanTilt.Status = e1.Text()
			}
		}
		if e1 := e.FindElement("Zoom"); e1 != nil {
			if e2 := e1.FindElement("Status"); e2 != nil {
				ptzStatus.MoveStatus.Zoom.Status = e2.Text()
			} else {
				ptzStatus.MoveStatus.Zoom.Status = e1.Text()
			}
		}
	}
	if e := status.FindElement("Error"); e != nil {
		ptzStatus.Error = e.Text()
	}
	if e := status.FindElement("UtcTime"); e != nil {
		ptzStatus.UtcTime, _ = getTimeFromXsdDateTime(e)
	}

	return &ptzStatus, nil
}

func (dev *Device) DecodeSetPreset(data []byte) (*string, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	token := doc.FindElement("./Envelope/Body/SetPresetResponse/PresetToken")
	if token == nil {
		return nil, fmt.Errorf("PresetToken element not found")
	}

	presetToken := token.Text()

	return &presetToken, nil
}

func (dev *Device) GetFault(data []byte) (bool, string, error) {
	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return false, "", err
	}

	fault := doc.FindElement("./Envelope/Body/Fault")
	if fault == nil {
		return false, "", nil
	}

	faultInfo := []string{}

	code := fault.FindElement("Code")
	if code != nil {
		if value := code.FindElement("Value"); value != nil {
			faultInfo = append(faultInfo, value.Text())
		}
		if subcode := code.FindElement("Subcode"); subcode != nil {
			if value := subcode.FindElement("Value"); value != nil {
				faultInfo = append(faultInfo, value.Text())
			}
			if subsubcode := subcode.FindElement("Subcode"); subsubcode != nil {
				if value := subsubcode.FindElement("Value"); value != nil {
					faultInfo = append(faultInfo, value.Text())
				}
			}
		}
	}

	reason := fault.FindElement("Reason")
	if reason != nil {
		if text := reason.FindElement("Text"); text != nil {
			faultInfo = append(faultInfo, text.Text())
		}
	}

	text := reason.FindElement("Detail")
	if text != nil {
		faultInfo = append(faultInfo, text.Text())
	}

	if len(faultInfo) == 0 {
		return false, "", nil
	}

	return true, strings.Join(faultInfo, ","), nil
}
