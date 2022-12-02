package elements

// The NonIndexableItemStatistic element contains a single statistic for an item that could not be indexed
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nonindexableitemstatistic
type NonIndexableItemStatistic struct {
	// The ErrorMessage element represents the reason for the validation error.
	ErrorMessage *ErrorMessage `xml:"m:ErrorMessage"`
	// The ItemCount element specifies the total number of items in a search result.
	ItemCount *ItemCount `xml:"ItemCount"`
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
}
