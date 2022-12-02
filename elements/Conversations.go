package elements

// The Conversations element contains an array of conversations that are returned in the FindConversation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversations-ex15websvcsotherref
type Conversations struct {
	// The Conversation element represents a single conversation.
	Conversation *ConversationConversationType `xml:"t:Conversation"`
}
