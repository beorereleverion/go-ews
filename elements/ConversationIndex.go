package elements

// The ConversationIndex element contains a binary ID that represents the thread to which this message belongs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationindex
import "encoding/xml"

type ConversationIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ConversationIndex) SetForMarshal() {
	C.XMLName.Local = "t:ConversationIndex"
}

func (C *ConversationIndex) GetSchema() *Schema {
	return &SchemaTypes
}
