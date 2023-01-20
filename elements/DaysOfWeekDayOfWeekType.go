package elements

// The DaysOfWeek (DayOfWeekType) element describes days of the week that are used in item recurrence patterns.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/daysofweek-dayofweektype
import "encoding/xml"

type DaysOfWeekDayOfWeekType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DaysOfWeekDayOfWeekType) SetForMarshal() {
	D.XMLName.Local = "t:DaysOfWeek"
}

func (D *DaysOfWeekDayOfWeekType) GetSchema() *Schema {
	return &SchemaTypes
}
