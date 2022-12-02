package elements

// The Attendees element specifies the recipients of an invitation to a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendees
type Attendees struct {
	// The EmailUser element specifies an email recipient.
	EmailUser *EmailUser `xml:"t:EmailUser"`
}
