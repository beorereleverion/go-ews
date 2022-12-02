package elements

// The AbsoluteYearlyRecurrence element represents a yearly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absoluteyearlyrecurrence
type AbsoluteYearlyRecurrence struct {
	// The DayOfMonth element describes the day in a month that a recurring item occurs.
	DayOfMonth *DayOfMonth `xml:"t:DayOfMonth"`
	// The Month element describes the month when a yearly recurring item occurs.
	Month *MonthItemRecurrence `xml:"t:Month"`
}
