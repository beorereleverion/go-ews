package elements

// The Conversation element represents a single conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversation-conversationtype
import "encoding/xml"

type ConversationConversationType struct {
	XMLName xml.Name

	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"Categories"`
	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"ConversationId"`
	// The ConversationTopic element represents the conversation topic.
	ConversationTopic *ConversationTopic `xml:"ConversationTopic"`
	// The FlagStatus element contains the aggregated flag status for conversation items in the current folder.
	FlagStatus *FlagStatus `xml:"FlagStatus"`
	// The GlobalCategories element contains the category list for all conversation items in a mailbox.
	GlobalCategories *GlobalCategories `xml:"GlobalCategories"`
	// The GlobalFlagStatus element contains the aggregated flag status for all conversation items in a mailbox.
	GlobalFlagStatus *GlobalFlagStatus `xml:"GlobalFlagStatus"`
	// The GlobalHasAttachments element contains a value that indicates whether at least one conversation item in a mailbox has an attachment.
	GlobalHasAttachments *GlobalHasAttachments `xml:"GlobalHasAttachments"`
	// The GlobalImportance element contains the aggregated importance for all conversation items in a mailbox.
	GlobalImportance *GlobalImportance `xml:"GlobalImportance"`
	// The GlobalItemClasses element contains a list of item classes that represents all the item classes of the conversation items in a mailbox.
	GlobalItemClasses *GlobalItemClasses `xml:"GlobalItemClasses"`
	// The GlobalItemIds element contains the collection of item identifiers for all conversation items in a mailbox.
	GlobalItemIds *GlobalItemIds `xml:"GlobalItemIds"`
	// The GlobalLastDeliveryTime element contains the delivery time of the message that was last received in this conversation across all folders in the mailbox.
	GlobalLastDeliveryTime *GlobalLastDeliveryTime `xml:"GlobalLastDeliveryTime"`
	// The GlobalMessageCount element contains the total number of conversation items in the mailbox.
	GlobalMessageCount *GlobalMessageCount `xml:"GlobalMessageCount"`
	// The GlobalSize element contains the size of the conversation calculated from the size of all conversation items in the mailbox.
	GlobalSize *GlobalSize `xml:"GlobalSize"`
	// The GlobalUniqueRecipients element contains the recipient list of a conversation aggregated across a mailbox.
	GlobalUniqueRecipients *GlobalUniqueRecipients `xml:"GlobalUniqueRecipients"`
	// The GlobalUniqueSender element contains a list of all the senders of conversation items in the mailbox.
	GlobalUniqueSenders *GlobalUniqueSenders `xml:"GlobalUniqueSenders"`
	// The GlobalUniqueUnreadSenders element specifies a list of all the people who have sent messages that are currently unread in this conversation across all folders in the mailbox.
	GlobalUniqueUnreadSenders *GlobalUniqueUnreadSenders `xml:"GlobalUniqueUnreadSenders"`
	// The GlobalUnreadCount element contains a count of all the unread conversation items in the mailbox.
	GlobalUnreadCount *GlobalUnreadCount `xml:"GlobalUnreadCount"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"HasAttachments"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"Importance"`
	// The ItemClasses element contains a list of item classes that represents all the item classes of the conversation items in the current folder.
	ItemClasses *ItemClassesArrayOfItemClassType `xml:"ItemClasses"`
	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"ItemIds"`
	// The LastDeliveryTime element contains the delivery time of the message that was last received in this conversation in the current folder.
	LastDeliveryTime *LastDeliveryTime `xml:"LastDeliveryTime"`
	// The MessageCount element contains the total number of conversation items in the current folder.
	MessageCount *MessageCount `xml:"MessageCount"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"Size"`
	// The UniqueRecipients element contains the recipient list of a conversation in a particular folder. This element is read-only.
	UniqueRecipients *UniqueRecipients `xml:"UniqueRecipients"`
	// The UniqueSenders element contains a list of all the senders of conversation items in the current folder. This element is read-only.
	UniqueSenders *UniqueSenders `xml:"UniqueSenders"`
	// The UniqueUnreadSenders element contains a list of all the people who have sent messages that are currently unread in this conversation in the current folder. This element is read-only.
	UniqueUnreadSenders *UniqueUnreadSenders `xml:"UniqueUnreadSenders"`
	// The UnreadCount element contains the count of unread items within a folder.
	UnreadCount *UnreadCount `xml:"UnreadCount"`
}

func (C *ConversationConversationType) SetForMarshal() {
	C.XMLName.Local = "t:Conversation"
}

func (C *ConversationConversationType) GetSchema() *Schema {
	return &SchemaTypes
}
