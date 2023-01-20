package elements

// The Title element represents the title of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/title
import "encoding/xml"

type Title struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *Title) SetForMarshal() {
	T.XMLName.Local = "t:Title"
}

func (T *Title) GetSchema() *Schema {
	return &SchemaTypes
}
