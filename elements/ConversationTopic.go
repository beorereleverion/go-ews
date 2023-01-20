package elements

// The ConversationTopic element represents the conversation topic.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationtopic
import "encoding/xml"

type ConversationTopic struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ConversationTopic) SetForMarshal() {
	C.XMLName.Local = "t:ConversationTopic"
}

func (C *ConversationTopic) GetSchema() *Schema {
	return &SchemaTypes
}
