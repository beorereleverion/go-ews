package elements

// The AssignCategories element represents the categories that are stamped on e-mail messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assigncategories
import "encoding/xml"

type AssignCategories struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (A *AssignCategories) SetForMarshal() {
	A.XMLName.Local = "m:AssignCategories"
}

func (A *AssignCategories) GetSchema() *Schema {
	return &SchemaMessages
}
