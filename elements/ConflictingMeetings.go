package elements

// The ConflictingMeetings element identifies all calendar items that conflict with a meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conflictingmeetings
type ConflictingMeetings struct {
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
}
