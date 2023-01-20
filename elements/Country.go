package elements

// The Country element identifies a country identifier in a postal address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/country
import "encoding/xml"

type Country struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *Country) SetForMarshal() {
	C.XMLName.Local = "t:Country"
}

func (C *Country) GetSchema() *Schema {
	return &SchemaTypes
}
