package elements

// The RootFolder element contains the results of a search of a single root folder during a FindItem operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootfolder-finditemresponsemessage
type RootFolderFindItemResponseMessage struct {
	// The Groups element contains a collection of groups that are found with the search and aggregation criteria that is identified in the FindItem operation request.
	Groups *Groups `xml:"t:Groups"`
	// The Items element contains an array of items.
	Items *Items `xml:"t:Items"`
	// Represents the next denominator to use for the next request when doing fractional paging.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// Indicates whether the current results contain the last item in the query, such that further paging is not needed.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when using an indexed paging view.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when using fraction page views.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Represents the total number of items that pass the restriction. In a grouped FindItem operation, the TotalItemsInView attribute returns the total number of items in the view plus the total number of groups.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}
