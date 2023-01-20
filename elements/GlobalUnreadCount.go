package elements

// The GlobalUnreadCount element contains a count of all the unread conversation items in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalunreadcount
import "encoding/xml"

type GlobalUnreadCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (G *GlobalUnreadCount) SetForMarshal() {
	G.XMLName.Local = "t:GlobalUnreadCount"
}

func (G *GlobalUnreadCount) GetSchema() *Schema {
	return &SchemaTypes
}
