package elements

// The Attendee element represents attendees and resources for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendee
type Attendee struct {
	// The LastResponseTime element represents the date and time of the latest response received.
	LastResponseTime *LastResponseTime `xml:"t:LastResponseTime"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
	// The ProposedEnd element specifies the proposed end time of a meeting.
	ProposedEnd *ProposedEnd `xml:"t:ProposedEnd"`
	// The ProposedStart element specifies the proposed start time of a meeting.
	ProposedStart *ProposedStart `xml:"t:ProposedStart"`
	// The ResponseType element represents the type of recipient response that is received for a meeting.
	ResponseType *ResponseType `xml:"t:ResponseType"`
}
