package elements

// The AbsoluteMonthlyRecurrence element represents a monthly recurrence pattern.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absolutemonthlyrecurrence
type AbsoluteMonthlyRecurrence struct {
	// The DayOfMonth element describes the day in a month that a recurring item occurs.
	DayOfMonth *DayOfMonth `xml:"t:DayOfMonth"`
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
