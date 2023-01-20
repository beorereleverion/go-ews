package elements

// The ImAddresses element represents a collection of instant messaging addresses for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddresses
import "encoding/xml"

type ImAddresses struct {
	XMLName xml.Name

	// The Entry element represents an instant messaging (IM) address for a contact.
	Entry *EntryIMAddress `xml:"Entry"`
}

func (I *ImAddresses) SetForMarshal() {
	I.XMLName.Local = "t:ImAddresses"
}

func (I *ImAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
