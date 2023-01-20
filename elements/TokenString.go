package elements

// The Token element contains a search refiner token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/token-string
import "encoding/xml"

type TokenString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TokenString) SetForMarshal() {
	T.XMLName.Local = "t:Token"
}

func (T *TokenString) GetSchema() *Schema {
	return &SchemaTypes
}
