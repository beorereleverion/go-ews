package elements

// The Token element contains encrypted data that represents the identification token for the shared data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/token
import "encoding/xml"

type Token struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *Token) SetForMarshal() {
	T.XMLName.Local = "t:Token"
}

func (T *Token) GetSchema() *Schema {
	return &SchemaTypes
}
