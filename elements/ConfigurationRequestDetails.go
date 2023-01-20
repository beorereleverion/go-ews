package elements

// The ConfigurationRequestDetails element contains client state for policy nudges. State information can include the rules that are installed and client component version.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/configurationrequestdetails
import "encoding/xml"

type ConfigurationRequestDetails struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ConfigurationRequestDetails) SetForMarshal() {
	C.XMLName.Local = "m:ConfigurationRequestDetails"
}

func (C *ConfigurationRequestDetails) GetSchema() *Schema {
	return &SchemaMessages
}
