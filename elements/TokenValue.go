package elements

// The TokenValue element specifies the encoded client access token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tokenvalue
import "encoding/xml"

type TokenValue struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TokenValue) SetForMarshal() {
	T.XMLName.Local = "t:TokenValue"
}

func (T *TokenValue) GetSchema() *Schema {
	return &SchemaTypes
}
