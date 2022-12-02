package elements

// The OptionalAttendees element represents attendees who are not required to attend a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/optionalattendees
type OptionalAttendees struct {
	// The Attendee element represents attendees and resources for a meeting.
	Attendee *Attendee `xml:"t:Attendee"`
}
