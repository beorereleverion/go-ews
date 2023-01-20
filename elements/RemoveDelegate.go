package elements

// The RemoveDelegate element defines a request to remove delegates from a mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removedelegate
import "encoding/xml"

type RemoveDelegate struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
	// The UserIds element contains an array of delegate users to get or remove from a principal's mailbox. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UserIds *UserIds `xml:"UserIds"`
}

func (R *RemoveDelegate) SetForMarshal() {
	R.XMLName.Local = "m:RemoveDelegate"
}

func (R *RemoveDelegate) GetSchema() *Schema {
	return &SchemaMessages
}
