package elements

// The TokenType element identifies the type of client access token that will be returned in the GetClientAccessToken response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokentype-clientaccesstokentype
import "encoding/xml"

type TokenTypeClientAccessTokenType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// CallerIdentity
	TokenTypeClientAccessTokenTypeCallerIdentity string = `CallerIdentity`
	// ExtensionCallback
	TokenTypeClientAccessTokenTypeExtensionCallback string = `ExtensionCallback`
	// ScopedToken
	TokenTypeClientAccessTokenTypeScopedToken string = `ScopedToken`
)

func (T *TokenTypeClientAccessTokenType) SetForMarshal() {
	T.XMLName.Local = "t:TokenType"
}

func (T *TokenTypeClientAccessTokenType) GetSchema() *Schema {
	return &SchemaTypes
}
