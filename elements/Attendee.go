package elements

// The Attendee element represents attendees and resources for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/attendee
import "encoding/xml"

type Attendee struct {
	XMLName xml.Name

	// The LastResponseTime element represents the date and time of the latest response received.
	LastResponseTime *LastResponseTime `xml:"LastResponseTime"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
	// The ProposedEnd element specifies the proposed end time of a meeting.
	ProposedEnd *ProposedEnd `xml:"ProposedEnd"`
	// The ProposedStart element specifies the proposed start time of a meeting.
	ProposedStart *ProposedStart `xml:"ProposedStart"`
	// The ResponseType element represents the type of recipient response that is received for a meeting.
	ResponseType *ResponseType `xml:"ResponseType"`
}

func (A *Attendee) SetForMarshal() {
	A.XMLName.Local = "t:Attendee"
}

func (A *Attendee) GetSchema() *Schema {
	return &SchemaTypes
}
