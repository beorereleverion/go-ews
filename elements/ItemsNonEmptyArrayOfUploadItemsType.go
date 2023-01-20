package elements

// The Items element contains an array of items to upload into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-nonemptyarrayofuploaditemstype
import "encoding/xml"

type ItemsNonEmptyArrayOfUploadItemsType struct {
	XMLName xml.Name

	// The Item element represents a single item to upload into a mailbox.
	Item *ItemUploadItemType `xml:"Item"`
}

func (I *ItemsNonEmptyArrayOfUploadItemsType) SetForMarshal() {
	I.XMLName.Local = "m:Items"
}

func (I *ItemsNonEmptyArrayOfUploadItemsType) GetSchema() *Schema {
	return &SchemaMessages
}
