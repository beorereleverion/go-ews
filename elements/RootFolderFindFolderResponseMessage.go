package elements

// The RootFolder element contains the results of a search of a single root folder during a FindFolder operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootfolder-findfolderresponsemessage
type RootFolderFindFolderResponseMessage struct {
	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"t:Folders"`
	// Represents the next denominator to use for the next request when doing fractional paging.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// Indicates whether the current results contain the last folder in the query, such that further paging is not needed.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when using an indexed paging view.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when using fractional page views.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Represents the total number of folders that pass the restriction.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}
