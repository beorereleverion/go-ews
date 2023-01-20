package elements

// The Watermark element represents an event bookmark in the mailbox event queue.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/watermark
import "encoding/xml"

type Watermark struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *Watermark) SetForMarshal() {
	W.XMLName.Local = "t:Watermark"
}

func (W *Watermark) GetSchema() *Schema {
	return &SchemaTypes
}
