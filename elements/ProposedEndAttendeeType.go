package elements

// The ProposedEnd (AttendeeType) element specifies an attendee's proposed end time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend-attendeetype
import "time"

import "encoding/xml"

type ProposedEndAttendeeType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedEndAttendeeType) SetForMarshal() {
	P.XMLName.Local = "t:ProposedEnd"
}

func (P *ProposedEndAttendeeType) GetSchema() *Schema {
	return &SchemaTypes
}
