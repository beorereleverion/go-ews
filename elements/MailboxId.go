package elements

// The MailboxId element specifies an identifier for the mailbox that is accessed by discovery search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxid
import "encoding/xml"

type MailboxId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MailboxId) SetForMarshal() {
	M.XMLName.Local = "t:MailboxId"
}

func (M *MailboxId) GetSchema() *Schema {
	return &SchemaTypes
}
