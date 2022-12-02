package elements

// The SuggestionArray element contains an array of meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionarray
type SuggestionArray struct {
	// The Suggestion element represents a single meeting suggestion.
	Suggestion *Suggestion `xml:"t:Suggestion"`
}
