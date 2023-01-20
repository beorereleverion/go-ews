package elements

// The IsRead element indicates whether a message has been read.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isread
import "encoding/xml"

type IsRead struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (I *IsRead) SetForMarshal() {
	I.XMLName.Local = "t:IsRead"
}

func (I *IsRead) GetSchema() *Schema {
	return &SchemaTypes
}
