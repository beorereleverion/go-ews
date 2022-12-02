package elements

// The ConversationActions element contains a collection of conversations and the actions to apply to them.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationactions
type ConversationActions struct {
	// The ConversationAction element contains a single action to be applied to a single conversation.
	ConversationAction *ConversationAction `xml:"t:ConversationAction"`
}
