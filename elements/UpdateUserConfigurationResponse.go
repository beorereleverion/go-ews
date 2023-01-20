package elements

// The UpdateUserConfigurationResponse element defines a response to a single UpdateUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateuserconfigurationresponse
import "encoding/xml"

type UpdateUserConfigurationResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (U *UpdateUserConfigurationResponse) SetForMarshal() {
	U.XMLName.Local = "m:UpdateUserConfigurationResponse"
}

func (U *UpdateUserConfigurationResponse) GetSchema() *Schema {
	return &SchemaMessages
}
