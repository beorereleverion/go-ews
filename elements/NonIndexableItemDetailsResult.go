package elements

// The NonIndexableItemDetailsResult element specifies the results of the GetNonIndexableItemDetails WSDL operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemdetailsresult
type NonIndexableItemDetailsResult struct {
	// The FailedMailboxes element specifies an array of mailboxes that failed on search.
	FailedMailboxes *FailedMailboxes `xml:"t:FailedMailboxes"`
	// The Items element contains an array of item details for non-indexable items.
	Items *ItemsArrayOfNonIndexableItemDetailsType `xml:"t:Items"`
}
