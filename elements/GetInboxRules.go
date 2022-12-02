package elements

// The GetInboxRules element defines a request to get the Inbox rules on a mailbox in the server store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getinboxrules
type GetInboxRules struct {
	// The MailboxSmtpAddress element represents the SMTP address of the user whose Inbox rules are to be retrieved or updated; or whose password expiration date is to be retrieved.
	MailboxSmtpAddress *MailboxSmtpAddress `xml:"t:MailboxSmtpAddress"`
}
