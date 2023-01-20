package elements

// The TokenType element specifies the type of token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokentype
import "encoding/xml"

type TokenType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// CallerIdentity
	TokenTypeCallerIdentity string = `CallerIdentity`
	// ExtensionCallback
	TokenTypeExtensionCallback string = `ExtensionCallback`
	// ScopedToken
	TokenTypeScopedToken string = `ScopedToken`
)

func (T *TokenType) SetForMarshal() {
	T.XMLName.Local = "t:TokenType"
}

func (T *TokenType) GetSchema() *Schema {
	return &SchemaTypes
}
