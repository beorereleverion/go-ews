package elements

// The DailyRegeneration element describes the frequency, in days, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dailyregeneration
type DailyRegeneration struct {
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
