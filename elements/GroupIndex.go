package elements

// The GroupIndex element represents the property value that is used to group items for the current group of items in a FindItem operation call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupindex
import "encoding/xml"

type GroupIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (G *GroupIndex) SetForMarshal() {
	G.XMLName.Local = "t:GroupIndex"
}

func (G *GroupIndex) GetSchema() *Schema {
	return &SchemaTypes
}
