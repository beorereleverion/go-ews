package elements

// The UploadItems element represents a request to upload items into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uploaditems
type UploadItems struct {
	// The Items element contains an array of items to upload into a mailbox.
	Items *ItemsNonEmptyArrayOfUploadItemsType `xml:"m:Items"`
}
