package elements

// The Items element contains an array of item details for non-indexable items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-arrayofnonindexableitemdetailstype
import "encoding/xml"

type ItemsArrayOfNonIndexableItemDetailsType struct {
	XMLName xml.Name

	// The NonIndexableItemDetail element specifies detail information about an item that cannot be indexed.
	NonIndexableItemDetail *NonIndexableItemDetail `xml:"NonIndexableItemDetail"`
}

func (I *ItemsArrayOfNonIndexableItemDetailsType) SetForMarshal() {
	I.XMLName.Local = "t:Items"
}

func (I *ItemsArrayOfNonIndexableItemDetailsType) GetSchema() *Schema {
	return &SchemaTypes
}
