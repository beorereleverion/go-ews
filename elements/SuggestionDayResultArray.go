package elements

// The SuggestionDayResultArray element contains an array of meeting suggestions organized by date.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestiondayresultarray
type SuggestionDayResultArray struct {
	// The SuggestionDayResult element represents a single day that contains suggested meeting times.
	SuggestionDayResult *SuggestionDayResult `xml:"t:SuggestionDayResult"`
}
