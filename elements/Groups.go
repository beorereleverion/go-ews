package elements

// The Groups element contains a collection of groups that are found with the search and aggregation criteria that is identified in the FindItem operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groups
import "encoding/xml"

type Groups struct {
	XMLName xml.Name

	// The GroupedItems element represents a collection of items that are the result of a grouped FindItem operation call.
	GroupedItems *GroupedItems `xml:"GroupedItems"`
}

func (G *Groups) SetForMarshal() {
	G.XMLName.Local = "t:Groups"
}

func (G *Groups) GetSchema() *Schema {
	return &SchemaTypes
}
