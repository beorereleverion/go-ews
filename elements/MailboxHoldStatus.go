package elements

// The MailboxHoldStatus element specifies the hold status of the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxholdstatus
type MailboxHoldStatus struct {
	// The AdditionalInfo element specifies additional information about the hold status of a mailbox.
	AdditionalInfo *AdditionalInfo `xml:"t:AdditionalInfo"`
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
	// The Status element specifies the hold status for a mailbox.
	Status *StatusHoldStatusType `xml:"t:Status"`
}
