package elements

// The ReceiveCopiesOfMeetingMessages element indicates whether a delegate receives copies of meeting-related messages that are addressed to the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivecopiesofmeetingmessages
import "encoding/xml"

type ReceiveCopiesOfMeetingMessages struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ReceiveCopiesOfMeetingMessagesfalse bool = false
	// true
	ReceiveCopiesOfMeetingMessagestrue bool = true
)

func (R *ReceiveCopiesOfMeetingMessages) SetForMarshal() {
	R.XMLName.Local = "t:ReceiveCopiesOfMeetingMessages"
}

func (R *ReceiveCopiesOfMeetingMessages) GetSchema() *Schema {
	return &SchemaTypes
}
