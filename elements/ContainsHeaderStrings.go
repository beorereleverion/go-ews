package elements

// The ContainsHeaderStrings element indicates the strings that must appear in the headers of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/containsheaderstrings
import "encoding/xml"

type ContainsHeaderStrings struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *ContainsHeaderStrings) SetForMarshal() {
	C.XMLName.Local = "m:ContainsHeaderStrings"
}

func (C *ContainsHeaderStrings) GetSchema() *Schema {
	return &SchemaMessages
}
