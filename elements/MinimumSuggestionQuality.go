package elements

// The MinimumSuggestionQuality element defines the quality of meeting suggestions to be returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/minimumsuggestionquality
type MinimumSuggestionQuality string

const (
	// 0% of the attendees have a conflict with the suggested meeting time.
	MinimumSuggestionQualityExcellent MinimumSuggestionQuality = `Excellent`
	// The percentage that is considered fair is set by using the GoodThreshold element.
	MinimumSuggestionQualityFair MinimumSuggestionQuality = `Fair`
	// The percentage that is considered good is set by using the GoodThreshold element.
	MinimumSuggestionQualityGood MinimumSuggestionQuality = `Good`
	// 50% or more of the attendees have a conflict with the suggested meeting time.
	MinimumSuggestionQualityPoor MinimumSuggestionQuality = `Poor`
)
