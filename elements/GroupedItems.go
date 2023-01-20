package elements

// The GroupedItems element represents a collection of items that are the result of a grouped FindItem operation call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupeditems
import "encoding/xml"

type GroupedItems struct {
	XMLName xml.Name

	// The GroupIndex element represents the property value that is used to group items for the current group of items in a FindItem operation call.
	GroupIndex *GroupIndex `xml:"GroupIndex"`
	// The Items element contains an array of items.
	Items *Items `xml:"Items"`
}

func (G *GroupedItems) SetForMarshal() {
	G.XMLName.Local = "t:GroupedItems"
}

func (G *GroupedItems) GetSchema() *Schema {
	return &SchemaTypes
}
