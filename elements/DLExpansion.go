package elements

// The DLExpansion element contains an array of mailboxes that are contained in a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dlexpansion
type DLExpansion struct {
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
	// Represents the next denominator to use for the next request when you are using fraction page views.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// Indicates whether the current results contain the last item in the query so that additional paging is not needed.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when you are using an indexed page view.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when you are using fraction page views.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Represents the total number of items in the view.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}
