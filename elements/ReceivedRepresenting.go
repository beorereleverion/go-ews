package elements

// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedrepresenting
import "encoding/xml"

type ReceivedRepresenting struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (R *ReceivedRepresenting) SetForMarshal() {
	R.XMLName.Local = "t:ReceivedRepresenting"
}

func (R *ReceivedRepresenting) GetSchema() *Schema {
	return &SchemaTypes
}
