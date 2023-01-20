package elements

// The Resolution element contains a single resolved entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolution
import "encoding/xml"

type Resolution struct {
	XMLName xml.Name

	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"Contact"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (R *Resolution) SetForMarshal() {
	R.XMLName.Local = "t:Resolution"
}

func (R *Resolution) GetSchema() *Schema {
	return &SchemaTypes
}
