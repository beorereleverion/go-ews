package elements

// The DaysOfWeek (DaysOfWeekType) element describes days of the week that are used in item recurrence patterns.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/daysofweek-daysofweektype
import "encoding/xml"

type DaysOfWeekDaysOfWeekType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DaysOfWeekDaysOfWeekType) SetForMarshal() {
	D.XMLName.Local = "t:DaysOfWeek"
}

func (D *DaysOfWeekDaysOfWeekType) GetSchema() *Schema {
	return &SchemaTypes
}
