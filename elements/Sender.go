package elements

// The Sender element identifies the sender of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sender
import "encoding/xml"

type Sender struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (S *Sender) SetForMarshal() {
	S.XMLName.Local = "t:Sender"
}

func (S *Sender) GetSchema() *Schema {
	return &SchemaTypes
}
