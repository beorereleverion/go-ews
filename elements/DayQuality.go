package elements

// The DayQuality element represents the quality of the day for containing quality suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dayquality
type DayQuality string

const (
	// Excellent
	DayQualityExcellent DayQuality = `Excellent`
	// Fair
	DayQualityFair DayQuality = `Fair`
	// Good
	DayQualityGood DayQuality = `Good`
	// Poor
	DayQualityPoor DayQuality = `Poor`
)
