package elements

// The ConversationActions element contains a collection of conversations and the actions to apply to them.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationactions
import "encoding/xml"

type ConversationActions struct {
	XMLName xml.Name

	// The ConversationAction element contains a single action to be applied to a single conversation.
	ConversationAction *ConversationAction `xml:"ConversationAction"`
}

func (C *ConversationActions) SetForMarshal() {
	C.XMLName.Local = "m:ConversationActions"
}

func (C *ConversationActions) GetSchema() *Schema {
	return &SchemaMessages
}
