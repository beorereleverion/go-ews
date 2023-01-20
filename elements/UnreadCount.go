package elements

// The UnreadCount element contains the count of unread items within a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unreadcount
import "encoding/xml"

type UnreadCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (U *UnreadCount) SetForMarshal() {
	U.XMLName.Local = "t:UnreadCount"
}

func (U *UnreadCount) GetSchema() *Schema {
	return &SchemaTypes
}
