package elements

// The GetUserAvailabilityResponse element is the root element that contains the properties that define user availability information or suggested meeting time information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailabilityresponse
type GetUserAvailabilityResponse struct {
	// The FreeBusyResponseArray element contains the requested users' availability information and the response status.
	FreeBusyResponseArray *FreeBusyResponseArray `xml:"m:FreeBusyResponseArray"`
	// The SuggestionsResponse element contains response status information and suggestion data for requested meeting suggestions.
	SuggestionsResponse *SuggestionsResponse `xml:"m:SuggestionsResponse"`
}
