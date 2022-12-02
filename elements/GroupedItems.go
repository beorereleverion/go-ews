package elements

// The GroupedItems element represents a collection of items that are the result of a grouped FindItem operation call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupeditems
type GroupedItems struct {
	// The GroupIndex element represents the property value that is used to group items for the current group of items in a FindItem operation call.
	GroupIndex *GroupIndex `xml:"t:GroupIndex"`
	// The Items element contains an array of items.
	Items *Items `xml:"t:Items"`
}
