package elements

// The Organizer element represents the organizer of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/organizer
import "encoding/xml"

type Organizer struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (O *Organizer) SetForMarshal() {
	O.XMLName.Local = "t:Organizer"
}

func (O *Organizer) GetSchema() *Schema {
	return &SchemaTypes
}
