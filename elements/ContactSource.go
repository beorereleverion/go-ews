package elements

// The ContactSource element describes whether the contact is located in the Exchange store or Active Directory Domain Services (AD DS).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactsource
import "encoding/xml"

type ContactSource struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ContactSource) SetForMarshal() {
	C.XMLName.Local = "t:ContactSource"
}

func (C *ContactSource) GetSchema() *Schema {
	return &SchemaTypes
}
