package elements

// The MarkAsRead element indicates whether messages are to be marked as read.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasread
import "encoding/xml"

type MarkAsRead struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	MarkAsReadfalse bool = false
	// true
	MarkAsReadtrue bool = true
)

func (M *MarkAsRead) SetForMarshal() {
	M.XMLName.Local = "m:MarkAsRead"
}

func (M *MarkAsRead) GetSchema() *Schema {
	return &SchemaMessages
}
