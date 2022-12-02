package elements

// The ConversationShape element identifies the property set to return in a FindConversation operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationshape
type ConversationShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"t:BaseShape"`
}
