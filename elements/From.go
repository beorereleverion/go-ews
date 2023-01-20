package elements

// The From element represents the address from which the message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/from
import "encoding/xml"

type From struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (F *From) SetForMarshal() {
	F.XMLName.Local = "t:From"
}

func (F *From) GetSchema() *Schema {
	return &SchemaTypes
}
