package elements

// The Token element specifies a client access token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/token-clientaccesstokentype
import "encoding/xml"

type TokenClientAccessTokenType struct {
	XMLName xml.Name

	// The ID element specifies the identifier of an app.
	ID *IDString `xml:"ID"`
	// The TTL element specifies the time, in minutes, that the token is valid.
	TTL *TTL `xml:"TTL"`
	// The TokenType element specifies the type of token.
	TokenType *TokenType `xml:"TokenType"`
	// The TokenValue element specifies the encoded client access token.
	TokenValue *TokenValue `xml:"TokenValue"`
}

func (T *TokenClientAccessTokenType) SetForMarshal() {
	T.XMLName.Local = "t:Token"
}

func (T *TokenClientAccessTokenType) GetSchema() *Schema {
	return &SchemaTypes
}
