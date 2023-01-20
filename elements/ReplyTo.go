package elements

// The ReplyTo element identifies an array of addresses to which replies should be sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/replyto
import "encoding/xml"

type ReplyTo struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (R *ReplyTo) SetForMarshal() {
	R.XMLName.Local = "t:ReplyTo"
}

func (R *ReplyTo) GetSchema() *Schema {
	return &SchemaTypes
}
