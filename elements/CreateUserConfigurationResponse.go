package elements

// The CreateUserConfigurationResponse element defines a response to a single CreateUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createuserconfigurationresponse
import "encoding/xml"

type CreateUserConfigurationResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateUserConfigurationResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateUserConfigurationResponse"
}

func (C *CreateUserConfigurationResponse) GetSchema() *Schema {
	return &SchemaMessages
}
