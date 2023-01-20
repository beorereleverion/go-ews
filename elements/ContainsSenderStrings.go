package elements

// The ContainsSenderStrings element indicates the strings that must appear in the From property of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/containssenderstrings
import "encoding/xml"

type ContainsSenderStrings struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *ContainsSenderStrings) SetForMarshal() {
	C.XMLName.Local = "m:ContainsSenderStrings"
}

func (C *ContainsSenderStrings) GetSchema() *Schema {
	return &SchemaMessages
}
