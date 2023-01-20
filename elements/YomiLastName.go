package elements

// The YomiLastName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese last name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomilastname
import "encoding/xml"

type YomiLastName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Y *YomiLastName) SetForMarshal() {
	Y.XMLName.Local = "t:YomiLastName"
}

func (Y *YomiLastName) GetSchema() *Schema {
	return &SchemaTypes
}
