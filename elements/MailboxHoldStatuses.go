package elements

// The MailboxHoldStatuses element specifies a list of one or more MailboxHoldStatus elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdstatuses
type MailboxHoldStatuses struct {
	// The MailboxHoldStatus element specifies the hold status of the mailbox.
	MailboxHoldStatus *MailboxHoldStatus `xml:"t:MailboxHoldStatus"`
}
