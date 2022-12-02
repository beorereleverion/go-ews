package elements

// The MonthlyRegeneration element describes the frequency, in months, of which task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/monthlyregeneration
type MonthlyRegeneration struct {
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
