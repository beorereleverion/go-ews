package elements

// The DirectReports element contains SMTP information that identifies the direct reports of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/directreports
import "encoding/xml"

type DirectReports struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (D *DirectReports) SetForMarshal() {
	D.XMLName.Local = "t:DirectReports"
}

func (D *DirectReports) GetSchema() *Schema {
	return &SchemaTypes
}
