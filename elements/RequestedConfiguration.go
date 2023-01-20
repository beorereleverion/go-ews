package elements

// The RequestedConfiguration element contains the requested service configurations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedconfiguration
import "encoding/xml"

type RequestedConfiguration struct {
	XMLName xml.Name

	// The ConfigurationName element specifies the requested service configurations by name.
	ConfigurationName *ConfigurationName `xml:"ConfigurationName"`
}

func (R *RequestedConfiguration) SetForMarshal() {
	R.XMLName.Local = "m:RequestedConfiguration"
}

func (R *RequestedConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
