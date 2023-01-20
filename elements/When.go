package elements

// The When element provides information about when a calendar or meeting item occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/when
import "encoding/xml"

type When struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *When) SetForMarshal() {
	W.XMLName.Local = "t:When"
}

func (W *When) GetSchema() *Schema {
	return &SchemaTypes
}
