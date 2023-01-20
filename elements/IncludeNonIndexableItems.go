package elements

// The IncludeNonIndexableItems element contains a Boolean value to indicate whether to include items that cannot be indexed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includenonindexableitems
import "encoding/xml"

type IncludeNonIndexableItems struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	IncludeNonIndexableItemsfalse bool = false
	// true
	IncludeNonIndexableItemstrue bool = true
)

func (I *IncludeNonIndexableItems) SetForMarshal() {
	I.XMLName.Local = "m:IncludeNonIndexableItems"
}

func (I *IncludeNonIndexableItems) GetSchema() *Schema {
	return &SchemaMessages
}
