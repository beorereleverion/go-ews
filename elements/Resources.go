package elements

// The Resources element represents a scheduled resource for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resources
type Resources struct {
	// The Attendee element represents attendees and resources for a meeting.
	Attendee *Attendee `xml:"t:Attendee"`
}
