package elements

// The SuggestionQuality element represents the quality of the suggested meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionquality
type SuggestionQuality string

const (
	// Excellent 100 percent of users and resources are available for the suggested meeting time.

	SuggestionQualityExcellent SuggestionQuality = `Excellent`
	// Fair The maximum percentage of users and resources available for a suggested meeting time is equal to the GoodThreshold element value plus 50. The minimum value for a Fair quality meeting time is 50 percent.

	SuggestionQualityFairFair SuggestionQuality = `FairFair`
	// Good The minimum percentage of users and resources available is equal to or greater than the GoodThreshold element value plus 50.

	SuggestionQualityGood SuggestionQuality = `Good`
	// Poor Less than 50 percent of the users and resources are available for the suggested meeting time.

	SuggestionQualityPoor SuggestionQuality = `Poor`
)
