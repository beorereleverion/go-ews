package elements

// The Entry element represents an instant messaging (IM) address for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entry-imaddress
import "encoding/xml"

type EntryIMAddress struct {
	XMLName xml.Name

	// Identifies the IM address.The following are the possible values for this attribute:- ImAddress1  - ImAddress2  - ImAddress3
	Key *string `xml:"Key,attr"`
}

func (E *EntryIMAddress) SetForMarshal() {
	E.XMLName.Local = "t:Entry"
}

func (E *EntryIMAddress) GetSchema() *Schema {
	return &SchemaTypes
}
