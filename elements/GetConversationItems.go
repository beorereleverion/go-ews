package elements

// The GetConversationItems element defines a request to get a set of items that are related by being in the same conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getconversationitems
import "encoding/xml"

type GetConversationItems struct {
	XMLName xml.Name

	// The Conversations element contains an array of conversations that are returned in the FindConversation response.
	Conversations *Conversations `xml:"Conversations"`
	// The FoldersToIgnore element identifies a list of folders that are ignored when getting items in a conversation. All conversation items in the ignored folders are not returned in a GetConversationItems response.
	FoldersToIgnore *FoldersToIgnore `xml:"FoldersToIgnore"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"ItemShape"`
	// The MailboxScope element identifies whether a search or fetch for a conversation should span either the primary mailbox, archive mailbox, or both the primary and archive mailbox.
	MailboxScope *MailboxScope `xml:"MailboxScope"`
	// The MaxItemsToReturn element identifies the maximum number of conversations items to return in a GetConversationItems response.
	MaxItemsToReturn *MaxItemsToReturn `xml:"MaxItemsToReturn"`
	// The SortOrder element specifies the sort order used for the result of a GetConversationItems request.
	SortOrder *SortOrderConversationNodeSortOrder `xml:"SortOrder"`
}

func (G *GetConversationItems) SetForMarshal() {
	G.XMLName.Local = "m:GetConversationItems"
}

func (G *GetConversationItems) GetSchema() *Schema {
	return &SchemaMessages
}
