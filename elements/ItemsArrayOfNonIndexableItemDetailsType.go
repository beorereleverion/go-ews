package elements

// The Items element contains an array of item details for non-indexable items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-arrayofnonindexableitemdetailstype
type ItemsArrayOfNonIndexableItemDetailsType struct {
	// The NonIndexableItemDetail element specifies detail information about an item that cannot be indexed.
	NonIndexableItemDetail *NonIndexableItemDetail `xml:"t:NonIndexableItemDetail"`
}
