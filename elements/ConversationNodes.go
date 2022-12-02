package elements

// The ConversationNodes element specifies a collection of conversation nodes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationnodes
type ConversationNodes struct {
	// The ConversationNode element specifies a node in a conversation.
	ConversationNode *ConversationNode `xml:"t:ConversationNode"`
}
