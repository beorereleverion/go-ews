package elements

// The MailboxFull element indicates whether the mailbox for the recipient is full.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxfull
type MailboxFull bool

const (
	// false
	MailboxFullfalse MailboxFull = false
	// true
	MailboxFulltrue MailboxFull = true
)
