package elements

// The MessageClassifications element represents the message classifications that must be stamped on incoming messages in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messageclassifications
import "encoding/xml"

type MessageClassifications struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (M *MessageClassifications) SetForMarshal() {
	M.XMLName.Local = "m:MessageClassifications"
}

func (M *MessageClassifications) GetSchema() *Schema {
	return &SchemaMessages
}
