package elements

// The MailboxCulture element indicates the culture to use when opening a mailbox. This element occurs in the SOAP header.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxculture
import "encoding/xml"

type MailboxCulture struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (M *MailboxCulture) SetForMarshal() {
	M.XMLName.Local = "t:MailboxCulture"
}

func (M *MailboxCulture) GetSchema() *Schema {
	return &SchemaTypes
}
