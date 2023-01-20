package elements

// The TokenRequest element specifies a single token request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokenrequest
import "encoding/xml"

type TokenRequest struct {
	XMLName xml.Name

	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"ID"`
	// The TokenType element identifies the type of client access token that will be returned in the GetClientAccessToken response.
	TokenType *TokenTypeClientAccessTokenType `xml:"TokenType"`
}

func (T *TokenRequest) SetForMarshal() {
	T.XMLName.Local = "t:TokenRequest"
}

func (T *TokenRequest) GetSchema() *Schema {
	return &SchemaTypes
}
