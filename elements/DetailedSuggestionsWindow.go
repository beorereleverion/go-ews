package elements

// The DetailedSuggestionsWindow element identifies the time span that is queried for detailed information about suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/detailedsuggestionswindow
type DetailedSuggestionsWindow struct {
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"t:EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"t:StartTime"`
}
