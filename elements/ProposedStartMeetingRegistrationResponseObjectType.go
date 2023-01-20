package elements

// The ProposedStart (MeetingRegistrationResponseObjectType) element specifies an attendee's proposed new start time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedstart-meetingregistrationresponseobjecttype
import "time"

import "encoding/xml"

type ProposedStartMeetingRegistrationResponseObjectType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedStartMeetingRegistrationResponseObjectType) SetForMarshal() {
	P.XMLName.Local = "t:ProposedStart"
}

func (P *ProposedStartMeetingRegistrationResponseObjectType) GetSchema() *Schema {
	return &SchemaTypes
}
