package elements

// The RelativeMonthlyRecurrence element describes a relative monthly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativemonthlyrecurrence
type RelativeMonthlyRecurrence struct {
	// The DayOfWeekIndex element describes which week in a month is used in a relative recurrence pattern.
	DayOfWeekIndex *DayOfWeekIndex `xml:"t:DayOfWeekIndex"`
	// The DaysOfWeek (DayOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDayOfWeekType `xml:"t:DaysOfWeek"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
