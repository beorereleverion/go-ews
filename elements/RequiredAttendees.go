package elements

// The RequiredAttendees element represents attendees that are required to attend a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requiredattendees
type RequiredAttendees struct {
	// The Attendee element represents attendees and resources for a meeting.
	Attendee *Attendee `xml:"t:Attendee"`
}
