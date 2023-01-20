package elements

// The GlobalUniqueUnreadSenders element specifies a list of all the people who have sent messages that are currently unread in this conversation across all folders in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaluniqueunreadsenders
import "encoding/xml"

type GlobalUniqueUnreadSenders struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (G *GlobalUniqueUnreadSenders) SetForMarshal() {
	G.XMLName.Local = "t:GlobalUniqueUnreadSenders"
}

func (G *GlobalUniqueUnreadSenders) GetSchema() *Schema {
	return &SchemaTypes
}
