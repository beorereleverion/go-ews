package elements

// The ResponseMessages element contains an array of service configuration response messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofserviceconfigurationresponsemessagetype
import "encoding/xml"

type ResponseMessagesArrayOfServiceConfigurationResponseMessageType struct {
	XMLName xml.Name

	// The ServiceConfigurationResponseMessageType element contains service configuration settings.
	ServiceConfigurationResponseMessageType *ServiceConfigurationResponseMessageType `xml:"ServiceConfigurationResponseMessageType"`
}

func (R *ResponseMessagesArrayOfServiceConfigurationResponseMessageType) SetForMarshal() {
	R.XMLName.Local = "m:ResponseMessages"
}

func (R *ResponseMessagesArrayOfServiceConfigurationResponseMessageType) GetSchema() *Schema {
	return &SchemaMessages
}
