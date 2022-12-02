package elements

// The AdjacentMeetings element identifies all calendar items that are adjacent to a meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adjacentmeetings
type AdjacentMeetings struct {
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
}
