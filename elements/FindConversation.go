package elements

// The FindConversation element defines a request to find conversations in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findconversation
import "encoding/xml"

type FindConversation struct {
	XMLName xml.Name

	// The ConversationShape element identifies the property set to return in a FindConversation operation response.
	ConversationShape *ConversationShape `xml:"ConversationShape"`
	// The IndexedPageItemView element describes how paged conversation or item information is returned for a FindItem operation or FindConversation operation request.
	IndexedPageItemView *IndexedPageItemView `xml:"IndexedPageItemView"`
	// The MailboxScope element identifies whether a search or fetch for a conversation should span either the primary mailbox, archive mailbox, or both the primary and archive mailbox.
	MailboxScope *MailboxScope `xml:"MailboxScope"`
	// The ParentFolderId element identifies the folder in which a new folder is created or the folder to search for the FindConversation operation.
	ParentFolderId *ParentFolderIdTargetFolderIdType `xml:"ParentFolderId"`
	// The QueryString element contains a mailbox query string based on Advanced Query Syntax (AQS).
	QueryString *QueryStringQueryStringType `xml:"QueryString"`
	// The SeekToConditionPageItemView element identifies the condition that is used to identify the end of a search, the starting index of a search, the maximum entries to return, and the search directions for a FindItem or FindConversation search.
	SeekToConditionPageItemView *SeekToConditionPageItemView `xml:"SeekToConditionPageItemView"`
	// The SortOrder element defines how items are sorted in a FindItem or FindConversation request.
	SortOrder *SortOrder `xml:"SortOrder"`
	// Identifies the types of sub-tree traversal. This attribute is optional.
	Traversal *string `xml:"Traversal,attr"`
	// Identifies the types view filters. This attribute is optional.
	ViewFilter *string `xml:"ViewFilter,attr"`
}

func (F *FindConversation) SetForMarshal() {
	F.XMLName.Local = "m:FindConversation"
}

func (F *FindConversation) GetSchema() *Schema {
	return &SchemaMessages
}
