package elements

// The SetHoldOnMailboxes element contains a SetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setholdonmailboxes
type SetHoldOnMailboxes struct {
	// The ActionType element indicates the type of action for the hold.
	ActionType *ActionTypeHoldActionType `xml:"t:ActionType"`
	// The Deduplication element indicates whether the search result should remove duplicate items.
	Deduplication *Deduplication `xml:"t:Deduplication"`
	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"t:HoldId"`
	// The InPlaceHoldIdentity element specifies the identity of a hold that preserves the mailbox items.
	InPlaceHoldIdentity *InPlaceHoldIdentity `xml:"m:InPlaceHoldIdentity"`
	// The IncludeNonIndexableItems element contains a Boolean value to indicate whether to include items that cannot be indexed.
	IncludeNonIndexableItems *IncludeNonIndexableItems `xml:"m:IncludeNonIndexableItems"`
	// The Language element contains the language used for the search query.
	Language *Language `xml:"m:Language"`
	// The Mailboxes element contains a list of mailboxes affected by the hold.
	Mailboxes *MailboxesArrayOfStringsType `xml:"m:Mailboxes"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"m:Query"`
}
