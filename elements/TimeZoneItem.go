package elements

// The TimeZone element provides a text description of a time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timezone-item
import "encoding/xml"

type TimeZoneItem struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TimeZoneItem) SetForMarshal() {
	T.XMLName.Local = "t:TimeZone"
}

func (T *TimeZoneItem) GetSchema() *Schema {
	return &SchemaTypes
}
