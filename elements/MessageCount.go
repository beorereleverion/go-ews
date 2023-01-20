package elements

// The MessageCount element contains the total number of conversation items in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagecount
import "encoding/xml"

type MessageCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MessageCount) SetForMarshal() {
	M.XMLName.Local = "t:MessageCount"
}

func (M *MessageCount) GetSchema() *Schema {
	return &SchemaTypes
}
