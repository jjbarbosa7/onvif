package analytics

import (
	"github.com/jjbarbosa7/onvif/xsd"
	"github.com/jjbarbosa7/onvif/xsd_onvif"
)

type GetSupportedRules struct {
	XMLName            string                   `xml:"tan:GetSupportedRules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type CreateRules struct {
	XMLName            string                   `xml:"tan:CreateRules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	Rule               xsd_onvif.Config         `xml:"tan:Rule"`
}

type DeleteRules struct {
	XMLName            string                   `xml:"tan:DeleteRules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	RuleName           xsd.String               `xml:"tan:RuleName"`
}

type GetRules struct {
	XMLName            string                   `xml:"tan:GetRules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetRuleOptions struct {
	XMLName            string                   `xml:"tan:GetRuleOptions"`
	RuleType           xsd.QName                `xml:"tan:RuleType"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type ModifyRules struct {
	XMLName            string                   `xml:"tan:ModifyRules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	Rule               xsd_onvif.Config         `xml:"tan:Rule"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"tan:GetServiceCapabilities"`
}

type GetSupportedAnalyticsModules struct {
	XMLName            string                   `xml:"tan:GetSupportedAnalyticsModules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetAnalyticsModuleOptions struct {
	XMLName            string                   `xml:"tan:GetAnalyticsModuleOptions"`
	Type               xsd.QName                `xml:"tan:Type"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type CreateAnalyticsModules struct {
	XMLName            string                   `xml:"tev:CreateAnalyticsModules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	AnalyticsModule    xsd_onvif.Config         `xml:"tan:AnalyticsModule"`
}

type DeleteAnalyticsModules struct {
	XMLName             string                   `xml:"tan:DeleteAnalyticsModules"`
	ConfigurationToken  xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	AnalyticsModuleName xsd.String               `xml:"tan:AnalyticsModuleName"`
}

type GetAnalyticsModules struct {
	XMLName            string                   `xml:"tan:GetAnalyticsModules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type ModifyAnalyticsModules struct {
	XMLName            string                   `xml:"tan:ModifyAnalyticsModules"`
	ConfigurationToken xsd_onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	AnalyticsModule    xsd_onvif.Config         `xml:"tan:AnalyticsModule"`
}
