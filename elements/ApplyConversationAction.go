package elements

// The ApplyConversationAction element defines a request to apply actions to items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applyconversationaction
type ApplyConversationAction struct {
	// The ConversationActions element contains a collection of conversations and the actions to apply to them.
	ConversationActions *ConversationActions `xml:"m:ConversationActions"`
}
