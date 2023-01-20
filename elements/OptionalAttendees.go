package elements

// The OptionalAttendees element represents attendees who are not required to attend a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/optionalattendees
import "encoding/xml"

type OptionalAttendees struct {
	XMLName xml.Name

	// The Attendee element represents attendees and resources for a meeting.
	Attendee *Attendee `xml:"Attendee"`
}

func (O *OptionalAttendees) SetForMarshal() {
	O.XMLName.Local = "t:OptionalAttendees"
}

func (O *OptionalAttendees) GetSchema() *Schema {
	return &SchemaTypes
}
