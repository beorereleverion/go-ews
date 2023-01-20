package elements

// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-nonemptyarrayoflegacydnstype
import "encoding/xml"

type MailboxesNonEmptyArrayOfLegacyDNsType struct {
	XMLName xml.Name

	// The LegacyDN element identifies a mailbox by its legacy distinguished name.
	LegacyDN *LegacyDN `xml:"LegacyDN"`
}

func (M *MailboxesNonEmptyArrayOfLegacyDNsType) SetForMarshal() {
	M.XMLName.Local = "m:Mailboxes"
}

func (M *MailboxesNonEmptyArrayOfLegacyDNsType) GetSchema() *Schema {
	return &SchemaMessages
}
