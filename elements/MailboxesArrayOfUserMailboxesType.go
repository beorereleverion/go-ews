package elements

// The Mailboxes element contains an array of mailboxes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-arrayofusermailboxestype
import "encoding/xml"

type MailboxesArrayOfUserMailboxesType struct {
	XMLName xml.Name

	// The UserMailbox element identifies a user mailbox.
	UserMailbox *UserMailbox `xml:"UserMailbox"`
}

func (M *MailboxesArrayOfUserMailboxesType) SetForMarshal() {
	M.XMLName.Local = "m:Mailboxes"
}

func (M *MailboxesArrayOfUserMailboxesType) GetSchema() *Schema {
	return &SchemaMessages
}
