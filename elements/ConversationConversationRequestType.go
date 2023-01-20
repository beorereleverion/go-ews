package elements

// Represents a single conversation returned in a GetConversationItems response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversation-conversationrequesttype
import "encoding/xml"

type ConversationConversationRequestType struct {
	XMLName xml.Name

	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"ConversationId"`
	// The SyncState element specifies the synchronization state of a conversation.
	SyncState *SyncStatebase64Binary `xml:"SyncState"`
}

func (C *ConversationConversationRequestType) SetForMarshal() {
	C.XMLName.Local = "t:Conversation"
}

func (C *ConversationConversationRequestType) GetSchema() *Schema {
	return &SchemaTypes
}
