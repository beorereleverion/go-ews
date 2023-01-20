package elements

// The PhysicalAddresses element contains a collection of physical addresses that are associated with a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/physicaladdresses
import "encoding/xml"

type PhysicalAddresses struct {
	XMLName xml.Name

	// The Entry element describes a single physical address for a contact item.
	Entry *EntryPhysicalAddress `xml:"Entry"`
}

func (P *PhysicalAddresses) SetForMarshal() {
	P.XMLName.Local = "t:PhysicalAddresses"
}

func (P *PhysicalAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
