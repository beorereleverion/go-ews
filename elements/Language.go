package elements

// The Language element contains the language used for the search query.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/language
import "encoding/xml"

type Language struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *Language) SetForMarshal() {
	L.XMLName.Local = "m:Language"
}

func (L *Language) GetSchema() *Schema {
	return &SchemaMessages
}
