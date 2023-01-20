package elements

// The InReplyTo element represents the identifier of the item to which this item is a reply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inreplyto
import "encoding/xml"

type InReplyTo struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (I *InReplyTo) SetForMarshal() {
	I.XMLName.Local = "t:InReplyTo"
}

func (I *InReplyTo) GetSchema() *Schema {
	return &SchemaTypes
}
