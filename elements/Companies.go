package elements

// The Companies element represents the collection of companies that are associated with a contact or task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/companies
import "encoding/xml"

type Companies struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *Companies) SetForMarshal() {
	C.XMLName.Local = "t:Companies"
}

func (C *Companies) GetSchema() *Schema {
	return &SchemaTypes
}
