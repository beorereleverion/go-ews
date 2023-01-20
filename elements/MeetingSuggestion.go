package elements

// The MeetingSuggestion element specifies a proposed meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingsuggestion
import "encoding/xml"

type MeetingSuggestion struct {
	XMLName xml.Name

	// The Attendees element specifies the recipients of an invitation to a meeting.
	Attendees *Attendees `xml:"Attendees"`
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"EndTime"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"Location"`
	// The MeetingString element specifies the name of the meeting as the result of entity extraction.
	MeetingString *MeetingString `xml:"MeetingString"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"StartTime"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
}

func (M *MeetingSuggestion) SetForMarshal() {
	M.XMLName.Local = "t:MeetingSuggestion"
}

func (M *MeetingSuggestion) GetSchema() *Schema {
	return &SchemaTypes
}
