package elements

// The Interval element defines the interval between two consecutive recurring items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/interval
import "encoding/xml"

type Interval struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (I *Interval) SetForMarshal() {
	I.XMLName.Local = "t:Interval"
}

func (I *Interval) GetSchema() *Schema {
	return &SchemaTypes
}
