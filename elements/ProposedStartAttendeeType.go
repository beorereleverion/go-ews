package elements

// The ProposedStart (AttendeeType) element specifies an attendee's proposed start time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart-attendeetype
import "time"

import "encoding/xml"

type ProposedStartAttendeeType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedStartAttendeeType) SetForMarshal() {
	P.XMLName.Local = "t:ProposedStart"
}

func (P *ProposedStartAttendeeType) GetSchema() *Schema {
	return &SchemaTypes
}
