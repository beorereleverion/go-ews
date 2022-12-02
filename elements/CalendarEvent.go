package elements

// The CalendarEvent element represents a unique calendar item occurrence.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarevent
type CalendarEvent struct {
	// The BusyType element represents the free/busy status set for a calendar event.
	BusyType *BusyType `xml:"t:BusyType"`
	// The CalendarEventDetails element provides additional information about a calendar event.
	CalendarEventDetails *CalendarEventDetails `xml:"t:CalendarEventDetails"`
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"t:EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"t:StartTime"`
}
