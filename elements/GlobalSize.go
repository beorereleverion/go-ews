package elements

// The GlobalSize element contains the size of the conversation calculated from the size of all conversation items in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalsize
import "encoding/xml"

type GlobalSize struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (G *GlobalSize) SetForMarshal() {
	G.XMLName.Local = "t:GlobalSize"
}

func (G *GlobalSize) GetSchema() *Schema {
	return &SchemaTypes
}
