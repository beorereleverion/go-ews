package elements

// The ManagerMailbox element contains SMTP information that identifies the mailbox of the contact's manager.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managermailbox
import "encoding/xml"

type ManagerMailbox struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (M *ManagerMailbox) SetForMarshal() {
	M.XMLName.Local = "t:ManagerMailbox"
}

func (M *ManagerMailbox) GetSchema() *Schema {
	return &SchemaTypes
}
