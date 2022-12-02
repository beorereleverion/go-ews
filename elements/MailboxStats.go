package elements

// The MailboxStats element specifies a list of one or more MailboxStat elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxstats
type MailboxStats struct {
	// The MailboxStat element specifies statistics for a mailbox searched by discovery search.
	MailboxStat *MailboxStat `xml:"t:MailboxStat"`
}
