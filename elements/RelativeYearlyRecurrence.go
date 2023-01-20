package elements

// The RelativeYearlyRecurrence element describes a relative yearly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativeyearlyrecurrence
import "encoding/xml"

type RelativeYearlyRecurrence struct {
	XMLName xml.Name

	// The DayOfWeekIndex element describes which week in a month is used in a relative recurrence pattern.
	DayOfWeekIndex *DayOfWeekIndex `xml:"DayOfWeekIndex"`
	// The DaysOfWeek (DayOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDayOfWeekType `xml:"DaysOfWeek"`
	// The Month element describes the month when a yearly recurring item occurs.
	Month *MonthItemRecurrence `xml:"Month"`
}

func (R *RelativeYearlyRecurrence) SetForMarshal() {
	R.XMLName.Local = "t:RelativeYearlyRecurrence"
}

func (R *RelativeYearlyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
