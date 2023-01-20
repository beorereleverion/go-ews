package elements

// The ConfigurationName element specifies the requested service configurations by name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/configurationname
import "encoding/xml"

type ConfigurationName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// MailTips
	ConfigurationNameMailTips string = `MailTips`
	// ProtectionRules
	ConfigurationNameProtectionRules string = `ProtectionRules`
	// UnifiedMessagingConfiguration
	ConfigurationNameUnifiedMessagingConfiguration string = `UnifiedMessagingConfiguration`
)

func (C *ConfigurationName) SetForMarshal() {
	C.XMLName.Local = "m:ConfigurationName"
}

func (C *ConfigurationName) GetSchema() *Schema {
	return &SchemaMessages
}
