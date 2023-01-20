package elements

// The ItemClasses element represents the item classes that must be stamped on incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemclasses
import "encoding/xml"

type ItemClasses struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (I *ItemClasses) SetForMarshal() {
	I.XMLName.Local = "t:ItemClasses"
}

func (I *ItemClasses) GetSchema() *Schema {
	return &SchemaTypes
}
