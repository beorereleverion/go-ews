package elements

// The DayOfWeek element represents the day of the week on which the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayofweek-timezone
import "encoding/xml"

type DayOfWeekTimeZone struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DayOfWeekTimeZone) SetForMarshal() {
	D.XMLName.Local = "t:DayOfWeek"
}

func (D *DayOfWeekTimeZone) GetSchema() *Schema {
	return &SchemaTypes
}
