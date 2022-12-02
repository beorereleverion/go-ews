package elements

// The YearlyRegeneration element describes the frequency, in years, in which a task is regenerated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yearlyregeneration
type YearlyRegeneration struct {
	// The Interval element defines the interval between two consecutive recurring items.
	Interval *Interval `xml:"t:Interval"`
}
