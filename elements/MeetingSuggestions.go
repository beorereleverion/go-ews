package elements

// The MeetingSuggestions element specifies an array of MeetingSuggestion elements that contain entity extraction results.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingsuggestions
import "encoding/xml"

type MeetingSuggestions struct {
	XMLName xml.Name

	// The MeetingSuggestion element specifies a proposed meeting.
	MeetingSuggestion *MeetingSuggestion `xml:"MeetingSuggestion"`
}

func (M *MeetingSuggestions) SetForMarshal() {
	M.XMLName.Local = "t:MeetingSuggestions"
}

func (M *MeetingSuggestions) GetSchema() *Schema {
	return &SchemaTypes
}
