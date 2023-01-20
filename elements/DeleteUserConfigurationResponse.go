package elements

// The DeleteUserConfigurationResponse element defines a response to a single DeleteUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteuserconfigurationresponse
import "encoding/xml"

type DeleteUserConfigurationResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (D *DeleteUserConfigurationResponse) SetForMarshal() {
	D.XMLName.Local = "m:DeleteUserConfigurationResponse"
}

func (D *DeleteUserConfigurationResponse) GetSchema() *Schema {
	return &SchemaMessages
}
