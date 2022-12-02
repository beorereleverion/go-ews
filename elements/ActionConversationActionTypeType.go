package elements

// The Action element contains the action to perform on the conversation specified by the ConversationId element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/action-conversationactiontypetype
type ActionConversationActionTypeType string

const (
	// AlwaysCategorize - The current items and new items in the conversation will automatically be set with the categories identified in the Categories element.

	ActionConversationActionTypeTypeAlwaysCategorize ActionConversationActionTypeType = `AlwaysCategorize`
	// AlwaysDelete - The current items and new items in the conversation will automatically be deleted. The deletion mode is set by the DeleteType element.

	ActionConversationActionTypeTypeAlwaysDelete ActionConversationActionTypeType = `AlwaysDelete`
	// AlwaysMove - The current items and new items in the conversation will automatically be moved to the folder identified by the DestinationFolderId element.

	ActionConversationActionTypeTypeAlwaysMove ActionConversationActionTypeType = `AlwaysMove`
	// Copy - The current items in the conversation will be copied to the folder identified by the DestinationFolderId element. Subsequent items in the conversation will not be copied.

	ActionConversationActionTypeTypeCopy ActionConversationActionTypeType = `Copy`
	// Delete - The current items in the conversation will be deleted. Subsequent items in the conversation will not be deleted. The deletion mode is set by the DeleteType element.

	ActionConversationActionTypeTypeDelete ActionConversationActionTypeType = `Delete`
	// Flag - The current items in the conversation will have a flag set as defined by the Flag element.

	ActionConversationActionTypeTypeFlag ActionConversationActionTypeType = `Flag`
	// Move - The current items in the conversation will be moved to the folder identified by the DestinationFolderId element. Subsequent items in the conversation will not be moved.

	ActionConversationActionTypeTypeMove ActionConversationActionTypeType = `Move`
	// SetReadState - The current items in the conversation will have their read state set. The read state is set by the IsRead element.

	ActionConversationActionTypeTypeSetReadState ActionConversationActionTypeType = `SetReadState`
)
