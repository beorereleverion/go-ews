package elements

// The ConversationNodes element specifies a collection of conversation nodes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationnodes
import "encoding/xml"

type ConversationNodes struct {
	XMLName xml.Name

	// The ConversationNode element specifies a node in a conversation.
	ConversationNode *ConversationNode `xml:"ConversationNode"`
}

func (C *ConversationNodes) SetForMarshal() {
	C.XMLName.Local = "t:ConversationNodes"
}

func (C *ConversationNodes) GetSchema() *Schema {
	return &SchemaTypes
}
