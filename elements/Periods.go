package elements

// The Periods element represents an array of periods that define the time offset at different stages of the time zone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/periods
type Periods struct {
	// The Period element defines the name, time offset, and unique identifier for a specific stage of the time zone.
	Period *Period `xml:"t:Period"`
}
