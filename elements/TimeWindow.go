package elements

// The TimeWindow element identifies the time span queried for the user availability information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timewindow
type TimeWindow struct {
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"t:EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"t:StartTime"`
}
