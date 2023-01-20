package elements

// The ExpandDL element defines a request to expand a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expanddl
import "encoding/xml"

type ExpandDL struct {
	XMLName xml.Name

	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (E *ExpandDL) SetForMarshal() {
	E.XMLName.Local = "m:ExpandDL"
}

func (E *ExpandDL) GetSchema() *Schema {
	return &SchemaMessages
}
