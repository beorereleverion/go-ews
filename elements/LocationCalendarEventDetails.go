package elements

// The Location element represents the location field of the calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/location-calendareventdetails
import "encoding/xml"

type LocationCalendarEventDetails struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LocationCalendarEventDetails) SetForMarshal() {
	L.XMLName.Local = "t:Location"
}

func (L *LocationCalendarEventDetails) GetSchema() *Schema {
	return &SchemaTypes
}
