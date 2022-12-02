package elements

// The ConversationNode element specifies a node in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationnode
type ConversationNode struct {
	// The InternetMessageId element represents the Internet message identifier of an item.
	InternetMessageId *InternetMessageId `xml:"t:InternetMessageId"`
	// The ItemIds element contains an array of item identifiers that identify the items to export from a mailbox.
	ItemIds *ItemIdsNonEmptyArrayOfItemIdsType `xml:"m:ItemIds"`
	// The ParentInternetMessageId element specifies the Internet message identifier of the parent message in a conversation.
	ParentInternetMessageId *ParentInternetMessageId `xml:"t:ParentInternetMessageId"`
}
