package elements

// The ApplyConversationAction element defines a request to apply actions to items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applyconversationaction
import "encoding/xml"

type ApplyConversationAction struct {
	XMLName xml.Name

	// The ConversationActions element contains a collection of conversations and the actions to apply to them.
	ConversationActions *ConversationActions `xml:"ConversationActions"`
}

func (A *ApplyConversationAction) SetForMarshal() {
	A.XMLName.Local = "m:ApplyConversationAction"
}

func (A *ApplyConversationAction) GetSchema() *Schema {
	return &SchemaMessages
}
