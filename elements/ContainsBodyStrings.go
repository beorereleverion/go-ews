package elements

// The ContainsBodyStrings element indicates the strings that must appear in the body of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/containsbodystrings
import "encoding/xml"

type ContainsBodyStrings struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *ContainsBodyStrings) SetForMarshal() {
	C.XMLName.Local = "m:ContainsBodyStrings"
}

func (C *ContainsBodyStrings) GetSchema() *Schema {
	return &SchemaMessages
}
