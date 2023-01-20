package elements

// The ReceivedBy element identifies the delegate in a delegate access scenario.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedby
import "encoding/xml"

type ReceivedBy struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (R *ReceivedBy) SetForMarshal() {
	R.XMLName.Local = "t:ReceivedBy"
}

func (R *ReceivedBy) GetSchema() *Schema {
	return &SchemaTypes
}
