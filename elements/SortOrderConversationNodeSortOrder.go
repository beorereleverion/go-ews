package elements

// The SortOrder element specifies the sort order used for the result of a GetConversationItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sortorder-conversationnodesortorder
type SortOrderConversationNodeSortOrder string

const (
	// DateOrderAscending
	SortOrderConversationNodeSortOrderDateOrderAscending SortOrderConversationNodeSortOrder = `DateOrderAscending`
	// DateOrderDescending
	SortOrderConversationNodeSortOrderDateOrderDescending SortOrderConversationNodeSortOrder = `DateOrderDescending`
	// TreeOrderAscending
	SortOrderConversationNodeSortOrderTreeOrderAscending SortOrderConversationNodeSortOrder = `TreeOrderAscending`
	// TreeOrderDescending
	SortOrderConversationNodeSortOrderTreeOrderDescending SortOrderConversationNodeSortOrder = `TreeOrderDescending`
)
