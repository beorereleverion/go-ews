package elements

// The EmailAddresses element represents a collection of e-mail addresses for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses
import "encoding/xml"

type EmailAddresses struct {
	XMLName xml.Name

	// The Entry element represents a single e-mail address for a contact.
	Entry *EntryEmailAddress `xml:"Entry"`
}

func (E *EmailAddresses) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddresses"
}

func (E *EmailAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
