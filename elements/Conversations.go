package elements

// The Conversations element contains an array of conversations that are returned in the FindConversation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversations-ex15websvcsotherref
import "encoding/xml"

type Conversations struct {
	XMLName xml.Name

	// The Conversation element represents a single conversation.
	Conversation *ConversationConversationType `xml:"Conversation"`
}

func (C *Conversations) SetForMarshal() {
	C.XMLName.Local = "m:Conversations"
}

func (C *Conversations) GetSchema() *Schema {
	return &SchemaMessages
}
