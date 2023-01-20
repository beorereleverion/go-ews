package elements

// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/importance
import "encoding/xml"

type Importance struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *Importance) SetForMarshal() {
	I.XMLName.Local = "t:Importance"
}

func (I *Importance) GetSchema() *Schema {
	return &SchemaTypes
}
