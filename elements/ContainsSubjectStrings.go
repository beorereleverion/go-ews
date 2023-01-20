package elements

// The ContainsSubjectStrings element indicates the strings that must appear in the subject of incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/containssubjectstrings
import "encoding/xml"

type ContainsSubjectStrings struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (C *ContainsSubjectStrings) SetForMarshal() {
	C.XMLName.Local = "m:ContainsSubjectStrings"
}

func (C *ContainsSubjectStrings) GetSchema() *Schema {
	return &SchemaMessages
}
