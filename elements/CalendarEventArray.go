package elements

// The CalendarEventArray element contains a set of unique calendar item occurrences that represent the requested user's availability.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendareventarray
type CalendarEventArray struct {
	// The CalendarEvent element represents a unique calendar item occurrence.
	CalendarEvent *CalendarEvent `xml:"t:CalendarEvent"`
}
