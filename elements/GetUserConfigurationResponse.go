package elements

// The GetUserConfigurationResponse element defines a response to a single GetUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserconfigurationresponse
import "encoding/xml"

type GetUserConfigurationResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetUserConfigurationResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetUserConfigurationResponse"
}

func (G *GetUserConfigurationResponse) GetSchema() *Schema {
	return &SchemaMessages
}
