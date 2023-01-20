package elements

// The Action element contains the action to perform on the conversation specified by the ConversationId element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/action-conversationactiontypetype
import "encoding/xml"

type ActionConversationActionTypeType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// AlwaysCategorize - The current items and new items in the conversation will automatically be set with the categories identified in the Categories element.

	ActionConversationActionTypeTypeAlwaysCategorize string = `AlwaysCategorize`
	// AlwaysDelete - The current items and new items in the conversation will automatically be deleted. The deletion mode is set by the DeleteType element.

	ActionConversationActionTypeTypeAlwaysDelete string = `AlwaysDelete`
	// AlwaysMove - The current items and new items in the conversation will automatically be moved to the folder identified by the DestinationFolderId element.

	ActionConversationActionTypeTypeAlwaysMove string = `AlwaysMove`
	// Copy - The current items in the conversation will be copied to the folder identified by the DestinationFolderId element. Subsequent items in the conversation will not be copied.

	ActionConversationActionTypeTypeCopy string = `Copy`
	// Delete - The current items in the conversation will be deleted. Subsequent items in the conversation will not be deleted. The deletion mode is set by the DeleteType element.

	ActionConversationActionTypeTypeDelete string = `Delete`
	// Flag - The current items in the conversation will have a flag set as defined by the Flag element.

	ActionConversationActionTypeTypeFlag string = `Flag`
	// Move - The current items in the conversation will be moved to the folder identified by the DestinationFolderId element. Subsequent items in the conversation will not be moved.

	ActionConversationActionTypeTypeMove string = `Move`
	// SetReadState - The current items in the conversation will have their read state set. The read state is set by the IsRead element.

	ActionConversationActionTypeTypeSetReadState string = `SetReadState`
)

func (A *ActionConversationActionTypeType) SetForMarshal() {
	A.XMLName.Local = "t:Action"
}

func (A *ActionConversationActionTypeType) GetSchema() *Schema {
	return &SchemaTypes
}
