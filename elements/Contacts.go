package elements

// The Contacts element contains a list of contacts that are associated with a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contacts-ex15websvcsotherref
import "encoding/xml"

type Contacts struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *Contacts) SetForMarshal() {
	C.XMLName.Local = "t:Contacts"
}

func (C *Contacts) GetSchema() *Schema {
	return &SchemaTypes
}
