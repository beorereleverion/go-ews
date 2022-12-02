package elements

// Represents a single conversation returned in a GetConversationItems response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversation-conversationrequesttype
type ConversationConversationRequestType struct {
	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"t:ConversationId"`
	// The SyncState element specifies the synchronization state of a conversation.
	SyncState *SyncStatebase64Binary `xml:"t:SyncState"`
}
