package elements

// The MessageTrackingSearchResults element contains a list of records that match the search criteria.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetrackingsearchresults
type MessageTrackingSearchResults struct {
	// The FindMessageTrackingReportResponse element contains the status and result of a single FindMessageTrackingReport operation request.
	FindMessageTrackingReportResponse *FindMessageTrackingReportResponse `xml:"m:FindMessageTrackingReportResponse"`
	// The MessageTrackingSearchResult element contains a single message result for a FindMessageTrackingReportResponse element.
	MessageTrackingSearchResult *MessageTrackingSearchResult `xml:"t:MessageTrackingSearchResult"`
}
