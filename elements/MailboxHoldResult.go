package elements

// The MailboxHoldResult element contains the result of the GetHoldOnMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdresult
type MailboxHoldResult struct {
	// The HoldId element contains the mailbox hold identifier.
	HoldId *HoldId `xml:"t:HoldId"`
	// The MailboxHoldStatuses element specifies a list of one or more MailboxHoldStatus elements.
	MailboxHoldStatuses *MailboxHoldStatuses `xml:"t:MailboxHoldStatuses"`
	// The Query element contains the search query for the hold.
	Query *Query `xml:"m:Query"`
}
