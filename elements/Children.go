package elements

// The Children element contains the names of a contact's children.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/children
import "encoding/xml"

type Children struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *Children) SetForMarshal() {
	C.XMLName.Local = "t:Children"
}

func (C *Children) GetSchema() *Schema {
	return &SchemaTypes
}
