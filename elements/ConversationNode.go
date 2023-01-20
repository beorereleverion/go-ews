package elements

// The ConversationNode element specifies a node in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationnode
import "encoding/xml"

type ConversationNode struct {
	XMLName xml.Name

	// The InternetMessageId element represents the Internet message identifier of an item.
	InternetMessageId *InternetMessageId `xml:"InternetMessageId"`
	// The ItemIds element contains an array of item identifiers that identify the items to export from a mailbox.
	ItemIds *ItemIdsNonEmptyArrayOfItemIdsType `xml:"ItemIds"`
	// The ParentInternetMessageId element specifies the Internet message identifier of the parent message in a conversation.
	ParentInternetMessageId *ParentInternetMessageId `xml:"ParentInternetMessageId"`
}

func (C *ConversationNode) SetForMarshal() {
	C.XMLName.Local = "t:ConversationNode"
}

func (C *ConversationNode) GetSchema() *Schema {
	return &SchemaTypes
}
