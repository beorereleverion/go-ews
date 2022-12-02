package elements

// The MeetingSuggestions element specifies an array of MeetingSuggestion elements that contain entity extraction results.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingsuggestions
type MeetingSuggestions struct {
	// The MeetingSuggestion element specifies a proposed meeting.
	MeetingSuggestion *MeetingSuggestion `xml:"t:MeetingSuggestion"`
}
