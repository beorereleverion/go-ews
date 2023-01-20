package elements

// The ID element represents the entry ID of the calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/id
import "encoding/xml"

type ID struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ID) SetForMarshal() {
	I.XMLName.Local = "t:ID"
}

func (I *ID) GetSchema() *Schema {
	return &SchemaTypes
}
