package elements

// The ConversationLastSyncTime element contains the date and time that a conversation was last synchronized. This element must be present when trying to delete all items in a conversation that were received up to the specified time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conversationlastsynctime
import "time"

import "encoding/xml"

type ConversationLastSyncTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *ConversationLastSyncTime) SetForMarshal() {
	C.XMLName.Local = "t:ConversationLastSyncTime"
}

func (C *ConversationLastSyncTime) GetSchema() *Schema {
	return &SchemaTypes
}
