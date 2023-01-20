package elements

// The ConversationIndex element specifies the location of a node in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationindex-string
import "encoding/xml"

type ConversationIndexstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ConversationIndexstring) SetForMarshal() {
	C.XMLName.Local = "t:ConversationIndex"
}

func (C *ConversationIndexstring) GetSchema() *Schema {
	return &SchemaTypes
}
