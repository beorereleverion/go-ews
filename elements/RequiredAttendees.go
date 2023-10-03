package elements

// The RequiredAttendees element represents attendees that are required to attend a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requiredattendees
import "encoding/xml"

type RequiredAttendees struct {
	XMLName xml.Name

	// The Attendee element represents attendees and resources for a meeting.
	Attendee []*Attendee `xml:"Attendee"`
}

func (R *RequiredAttendees) SetForMarshal() {
	R.XMLName.Local = "t:RequiredAttendees"
}

func (R *RequiredAttendees) GetSchema() *Schema {
	return &SchemaTypes
}
