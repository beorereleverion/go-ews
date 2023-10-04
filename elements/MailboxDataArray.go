package elements

// The MailboxDataArray element contains a list of mailboxes to query for availability information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxdataarray
import "encoding/xml"

type MailboxDataArray struct {
	XMLName xml.Name

	// The MailboxData element represents an individual mailbox user and options for the type of data to be returned about the mailbox user.
	MailboxData []*MailboxData `xml:"MailboxData"`
}

func (M *MailboxDataArray) SetForMarshal() {
	M.XMLName.Local = "m:MailboxDataArray"
}

func (M *MailboxDataArray) GetSchema() *Schema {
	return &SchemaMessages
}
