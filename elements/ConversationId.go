package elements

// The ConversationId element contains the identifier of an item or conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationid
import "encoding/xml"

type ConversationId struct {
	XMLName xml.Name

	// Identifies a specific version of an item. A ChangeKey is required for the following scenarios:  - The UpdateItem element requires a ChangeKey if the ConflictResolution attribute is set to AutoResolve. AutoResolve is a default value. If the ChangeKey attribute is not included, the response will return a ResponseCode value equal to ErrorChangeKeyRequired.- SendItem, DeleteItem, and DeleteFolder elements require a ChangeKey to test whether the attempted operation will act upon the most recent version of an item. If the ChangeKey attribute is not included in the ItemId or if the ChangeKey is empty, the response will return a ResponseCode value equal to ErrorStaleObject.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a specific item in the Exchange store.
	Id *string `xml:"Id,attr"`
}

func (C *ConversationId) SetForMarshal() {
	C.XMLName.Local = "t:ConversationId"
}

func (C *ConversationId) GetSchema() *Schema {
	return &SchemaTypes
}
