package elements

// The Recipients element represents a collection of recipients that receive a copy of the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipients-arrayofrecipientstype
import "encoding/xml"

type RecipientsArrayOfRecipientsType struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (R *RecipientsArrayOfRecipientsType) SetForMarshal() {
	R.XMLName.Local = "t:Recipients"
}

func (R *RecipientsArrayOfRecipientsType) GetSchema() *Schema {
	return &SchemaTypes
}
