package elements

// The MailboxFull element indicates whether the mailbox for the recipient is full.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxfull
import "encoding/xml"

type MailboxFull struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	MailboxFullfalse bool = false
	// true
	MailboxFulltrue bool = true
)

func (M *MailboxFull) SetForMarshal() {
	M.XMLName.Local = "t:MailboxFull"
}

func (M *MailboxFull) GetSchema() *Schema {
	return &SchemaTypes
}
