package elements

// The Suggestion element represents a single meeting suggestion.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestion
import "encoding/xml"

type Suggestion struct {
	XMLName xml.Name

	// The AttendeeConflictDataArray element contains an array of conflict data for queried attendees identified in the GetUserAvailability operation.
	AttendeeConflictDataArray *AttendeeConflictDataArray `xml:"AttendeeConflictDataArray"`
	// The IsWorkTime element represents whether the suggested meeting time occurs during scheduled work hours.
	IsWorkTime *IsWorkTime `xml:"IsWorkTime"`
	// The MeetingTime element represents a suggested meeting time.
	MeetingTime *MeetingTime `xml:"MeetingTime"`
	// The SuggestionQuality element represents the quality of the suggested meeting time.
	SuggestionQuality *SuggestionQuality `xml:"SuggestionQuality"`
}

func (S *Suggestion) SetForMarshal() {
	S.XMLName.Local = "t:Suggestion"
}

func (S *Suggestion) GetSchema() *Schema {
	return &SchemaTypes
}
