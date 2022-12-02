package elements

// The SuggestionsResponse element contains response status information and suggestion data for requested meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionsresponse
type SuggestionsResponse struct {
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"m:ResponseMessage"`
	// The SuggestionDayResultArray element contains an array of meeting suggestions organized by date.
	SuggestionDayResultArray *SuggestionDayResultArray `xml:"m:SuggestionDayResultArray"`
}
