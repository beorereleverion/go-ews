package elements

// The RelativeYearlyRecurrence element describes a relative yearly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativeyearlyrecurrence
type RelativeYearlyRecurrence struct {
	// The DayOfWeekIndex element describes which week in a month is used in a relative recurrence pattern.
	DayOfWeekIndex *DayOfWeekIndex `xml:"t:DayOfWeekIndex"`
	// The DaysOfWeek (DayOfWeekType) element describes days of the week that are used in item recurrence patterns.
	DaysOfWeek *DaysOfWeekDayOfWeekType `xml:"t:DaysOfWeek"`
	// The Month element describes the month when a yearly recurring item occurs.
	Month *MonthItemRecurrence `xml:"t:Month"`
}
