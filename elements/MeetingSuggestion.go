package elements

// The MeetingSuggestion element specifies a proposed meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingsuggestion
type MeetingSuggestion struct {
	// The Attendees element specifies the recipients of an invitation to a meeting.
	Attendees *Attendees `xml:"t:Attendees"`
	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"t:EndTime"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"t:Location"`
	// The MeetingString element specifies the name of the meeting as the result of entity extraction.
	MeetingString *MeetingString `xml:"t:MeetingString"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"t:StartTime"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"t:Subject"`
}
