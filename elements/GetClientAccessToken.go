package elements

// The GetClientAccessToken element contains a request to get a client access token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientaccesstoken
import "encoding/xml"

type GetClientAccessToken struct {
	XMLName xml.Name

	// The TokenRequests element contains an array of token requests.
	TokenRequests *TokenRequests `xml:"TokenRequests"`
}

func (G *GetClientAccessToken) SetForMarshal() {
	G.XMLName.Local = "m:GetClientAccessToken"
}

func (G *GetClientAccessToken) GetSchema() *Schema {
	return &SchemaMessages
}
