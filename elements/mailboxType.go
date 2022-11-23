package elements

type MailboxType string

const (
	// Represents a mail-enabled Active Directory object.
	MailboxTypeMailbox MailboxType = `Mailbox`
	// Represents a public distribution list.
	MailboxTypePublicDL MailboxType = `PublicDL`
	// Represents a private distribution list in a user's mailbox.
	MailboxTypePrivateDL MailboxType = `PrivateDL`
	// Represents a contact in a user's mailbox.
	MailboxTypeContact MailboxType = `Contact`
	// Represents a public folder.
	MailboxTypePublicFolder MailboxType = `PublicFolder`
	// Represents an unknown type of mailbox.
	MailboxTypeUnknown MailboxType = `Unknown`
	// Represents a one-off member of a personal distribution list.
	MailboxTypeOneOff MailboxType = `OneOff`
	// Represents a group mailbox.
	MailboxTypeGroupMailbox MailboxType = `GroupMailbox`
)
