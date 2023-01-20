package elements

// The MailboxScope element identifies whether a search or fetch for a conversation should span either the primary mailbox, archive mailbox, or both the primary and archive mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxscope
import "encoding/xml"

type MailboxScope struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// All
	MailboxScopeAll string = `All`
	// ArchiveOnly
	MailboxScopeArchiveOnly string = `ArchiveOnly`
	// PrimaryOnly
	MailboxScopePrimaryOnly string = `PrimaryOnly`
)

func (M *MailboxScope) SetForMarshal() {
	M.XMLName.Local = "m:MailboxScope"
}

func (M *MailboxScope) GetSchema() *Schema {
	return &SchemaMessages
}
