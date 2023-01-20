package elements

// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bccrecipients
import "encoding/xml"

type BccRecipients struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (B *BccRecipients) SetForMarshal() {
	B.XMLName.Local = "t:BccRecipients"
}

func (B *BccRecipients) GetSchema() *Schema {
	return &SchemaTypes
}
