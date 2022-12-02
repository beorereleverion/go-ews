package elements

// The DailyRecurrence element describes the frequency, in days, in which a calendar item or a task recurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dailyrecurrence
type DailyRecurrence struct {
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
