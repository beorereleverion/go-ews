package elements

// The GetClientAccessTokenResponse element contains the response to a GetClientAccessToken operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstokenresponse
import "encoding/xml"

type GetClientAccessTokenResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetClientAccessTokenResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetClientAccessTokenResponse"
}

func (G *GetClientAccessTokenResponse) GetSchema() *Schema {
	return &SchemaMessages
}
