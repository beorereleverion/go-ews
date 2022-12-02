package elements

// The WeeklyRegeneration element describes the frequency, in weeks, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weeklyregeneration
type WeeklyRegeneration struct {
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
