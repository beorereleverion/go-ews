package elements

// The YomiFirstName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese first name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomifirstname
import "encoding/xml"

type YomiFirstName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Y *YomiFirstName) SetForMarshal() {
	Y.XMLName.Local = "t:YomiFirstName"
}

func (Y *YomiFirstName) GetSchema() *Schema {
	return &SchemaTypes
}
