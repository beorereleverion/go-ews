package elements

// The RelativeMonthlyRecurrence element describes a relative monthly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativemonthlyrecurrence
import "encoding/xml"

type RelativeMonthlyRecurrence struct {
	XMLName xml.Name

	// The DayOfWeekIndex element describes which week in a month is used in a relative recurrence pattern.
	DayOfWeekIndex *DayOfWeekIndex `xml:"DayOfWeekIndex"`
	// The DaysOfWeek (DayOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDayOfWeekType `xml:"DaysOfWeek"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (R *RelativeMonthlyRecurrence) SetForMarshal() {
	R.XMLName.Local = "t:RelativeMonthlyRecurrence"
}

func (R *RelativeMonthlyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
