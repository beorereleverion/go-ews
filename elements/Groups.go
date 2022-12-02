package elements

// The Groups element contains a collection of groups that are found with the search and aggregation criteria that is identified in the FindItem operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groups
type Groups struct {
	// The GroupedItems element represents a collection of items that are the result of a grouped FindItem operation call.
	GroupedItems *GroupedItems `xml:"t:GroupedItems"`
}
