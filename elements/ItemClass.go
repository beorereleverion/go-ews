package elements

// The ItemClass element represents the message class of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemclass
import "encoding/xml"

type ItemClass struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *ItemClass) SetForMarshal() {
	I.XMLName.Local = "t:ItemClass"
}

func (I *ItemClass) GetSchema() *Schema {
	return &SchemaTypes
}
