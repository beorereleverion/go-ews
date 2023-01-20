package elements

// The DayOrder element represents the nth occurrence of the day specified in the DayOfWeek (TimeZone) element that represents the date of transition from and to standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayorder
import "encoding/xml"

type DayOrder struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (D *DayOrder) SetForMarshal() {
	D.XMLName.Local = "t:DayOrder"
}

func (D *DayOrder) GetSchema() *Schema {
	return &SchemaTypes
}
