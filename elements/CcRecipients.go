package elements

// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ccrecipients
import "encoding/xml"

type CcRecipients struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (C *CcRecipients) SetForMarshal() {
	C.XMLName.Local = "t:CcRecipients"
}

func (C *CcRecipients) GetSchema() *Schema {
	return &SchemaTypes
}
