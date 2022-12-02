package elements

// The FailedMailboxes element specifies an array of mailboxes that failed on search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/failedmailboxes
type FailedMailboxes struct {
	// The FailedMailbox element specifies the error message for a mailbox that failed on search.
	FailedMailbox *FailedMailbox `xml:"t:FailedMailbox"`
}
