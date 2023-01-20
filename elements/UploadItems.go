package elements

// The UploadItems element represents a request to upload items into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uploaditems
import "encoding/xml"

type UploadItems struct {
	XMLName xml.Name

	// The Items element contains an array of items to upload into a mailbox.
	Items *ItemsNonEmptyArrayOfUploadItemsType `xml:"Items"`
}

func (U *UploadItems) SetForMarshal() {
	U.XMLName.Local = "m:UploadItems"
}

func (U *UploadItems) GetSchema() *Schema {
	return &SchemaMessages
}
