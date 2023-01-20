package elements

// The Mailboxes element contains a list of mailboxes affected by the hold.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxes-arrayofstringstype
import "encoding/xml"

type MailboxesArrayOfStringsType struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (M *MailboxesArrayOfStringsType) SetForMarshal() {
	M.XMLName.Local = "m:Mailboxes"
}

func (M *MailboxesArrayOfStringsType) GetSchema() *Schema {
	return &SchemaMessages
}
