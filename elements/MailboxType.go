package elements

// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxtype
type MailboxType string

const (
	// Represents a contact in a user's mailbox.
	MailboxTypeContact MailboxType = `Contact`
	// Represents a group mailbox.
	MailboxTypeGroupMailbox MailboxType = `GroupMailbox`
	// Represents a mail-enabled Active Directory object.
	MailboxTypeMailbox MailboxType = `Mailbox`
	// Represents a one-off member of a personal distribution list.
	MailboxTypeOneOff MailboxType = `OneOff`
	// Represents a private distribution list in a user's mailbox.
	MailboxTypePrivateDL MailboxType = `PrivateDL`
	// Represents a public distribution list.
	MailboxTypePublicDL MailboxType = `PublicDL`
	// Represents a public folder.
	MailboxTypePublicFolder MailboxType = `PublicFolder`
	// Represents an unknown type of mailbox.
	MailboxTypeUnknown MailboxType = `Unknown`
)
