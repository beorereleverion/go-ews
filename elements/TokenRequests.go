package elements

// The TokenRequests element contains an array of token requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokenrequests
import "encoding/xml"

type TokenRequests struct {
	XMLName xml.Name

	// The TokenRequest element specifies a single token request.
	TokenRequest *TokenRequest `xml:"TokenRequest"`
}

func (T *TokenRequests) SetForMarshal() {
	T.XMLName.Local = "m:TokenRequests"
}

func (T *TokenRequests) GetSchema() *Schema {
	return &SchemaMessages
}
