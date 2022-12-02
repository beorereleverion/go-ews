package elements

// The WithinDateRange element specifies the date range within which incoming messages have to have been received in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/withindaterange
type WithinDateRange struct {
	// The EndDateTime element specifies the end date and time for a rule or a search.
	EndDateTime *EndDateTime `xml:"m:EndDateTime"`
	// The StartDateTime element specifies the start date and time for a rule or a search.
	StartDateTime *StartDateTime `xml:"m:StartDateTime"`
}
