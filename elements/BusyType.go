package elements

// The BusyType element represents the free/busy status set for a calendar event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/busytype
import "encoding/xml"

type BusyType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (B *BusyType) SetForMarshal() {
	B.XMLName.Local = "t:BusyType"
}

func (B *BusyType) GetSchema() *Schema {
	return &SchemaTypes
}
