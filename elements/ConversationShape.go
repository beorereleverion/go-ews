package elements

// The ConversationShape element identifies the property set to return in a FindConversation operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationshape
import "encoding/xml"

type ConversationShape struct {
	XMLName xml.Name

	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"AdditionalProperties"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape *BaseShape `xml:"BaseShape"`
}

func (C *ConversationShape) SetForMarshal() {
	C.XMLName.Local = "t:ConversationShape"
}

func (C *ConversationShape) GetSchema() *Schema {
	return &SchemaTypes
}
