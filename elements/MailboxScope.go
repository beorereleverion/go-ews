package elements

// The MailboxScope element identifies whether a search or fetch for a conversation should span either the primary mailbox, archive mailbox, or both the primary and archive mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxscope
type MailboxScope string

const (
	// All
	MailboxScopeAll MailboxScope = `All`
	// ArchiveOnly
	MailboxScopeArchiveOnly MailboxScope = `ArchiveOnly`
	// PrimaryOnly
	MailboxScopePrimaryOnly MailboxScope = `PrimaryOnly`
)
