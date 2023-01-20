package elements

// Conversation (ConversationResponseType) represents a single conversation returned in a GetConversationItems response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversation-conversationresponsetype
import "encoding/xml"

type ConversationConversationResponseType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
