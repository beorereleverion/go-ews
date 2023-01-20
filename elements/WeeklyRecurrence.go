package elements

// The WeeklyRecurrence element describes a weekly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weeklyrecurrence
import "encoding/xml"

type WeeklyRecurrence struct {
	XMLName xml.Name

	// The DaysOfWeek (DaysOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDaysOfWeekType `xml:"DaysOfWeek"`
	// The FirstDayOfWeek element specifies the first day of the week.
	FirstDayOfWeek *FirstDayOfWeek `xml:"FirstDayOfWeek"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"Interval"`
}

func (W *WeeklyRecurrence) SetForMarshal() {
	W.XMLName.Local = "t:WeeklyRecurrence"
}

func (W *WeeklyRecurrence) GetSchema() *Schema {
	return &SchemaTypes
}
