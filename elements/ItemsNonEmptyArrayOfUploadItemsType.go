package elements

// The Items element contains an array of items to upload into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-nonemptyarrayofuploaditemstype
type ItemsNonEmptyArrayOfUploadItemsType struct {
	// The Item element represents a single item to upload into a mailbox.
	Item *ItemUploadItemType `xml:"t:Item"`
}
