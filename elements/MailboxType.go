package elements

// The MailboxType element represents the type of mailbox that is represented by the e-mail address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxtype
import "encoding/xml"

type MailboxType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Represents a contact in a user's mailbox.
	MailboxTypeContact string = `Contact`
	// Represents a group mailbox.
	MailboxTypeGroupMailbox string = `GroupMailbox`
	// Represents a mail-enabled Active Directory object.
	MailboxTypeMailbox string = `Mailbox`
	// Represents a one-off member of a personal distribution list.
	MailboxTypeOneOff string = `OneOff`
	// Represents a private distribution list in a user's mailbox.
	MailboxTypePrivateDL string = `PrivateDL`
	// Represents a public distribution list.
	MailboxTypePublicDL string = `PublicDL`
	// Represents a public folder.
	MailboxTypePublicFolder string = `PublicFolder`
	// Represents an unknown type of mailbox.
	MailboxTypeUnknown string = `Unknown`
)

func (M *MailboxType) SetForMarshal() {
	M.XMLName.Local = "t:MailboxType"
}

func (M *MailboxType) GetSchema() *Schema {
	return &SchemaTypes
}
