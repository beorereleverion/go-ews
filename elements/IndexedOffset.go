package elements

// The IndexedOffset element indicates the index offset for a paged FindConversation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedoffset
import "encoding/xml"

type IndexedOffset struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (I *IndexedOffset) SetForMarshal() {
	I.XMLName.Local = "m:IndexedOffset"
}

func (I *IndexedOffset) GetSchema() *Schema {
	return &SchemaMessages
}
