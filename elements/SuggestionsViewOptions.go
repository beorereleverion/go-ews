package elements

// The SuggestionsViewOptions element contains the options for obtaining meeting suggestion information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionsviewoptions
import "encoding/xml"

type SuggestionsViewOptions struct {
	XMLName xml.Name

	// The CurrentMeetingTime element represents the start time of a meeting that you want to update with a meeting time proposed by a meeting attendee.
	CurrentMeetingTime *CurrentMeetingTime `xml:"CurrentMeetingTime"`
	// The DetailedSuggestionsWindow element identifies the time span that is queried for detailed information about suggested meeting times.
	DetailedSuggestionsWindow *DetailedSuggestionsWindow `xml:"DetailedSuggestionsWindow"`
	// The GlobalObjectId element is not used.
	GlobalObjectId *GlobalObjectId `xml:"GlobalObjectId"`
	// The GoodThreshold element specifies the percentage of attendees that must have the time period open in order for the time period to qualify as a Good suggested meeting time.
	GoodThreshold *GoodThreshold `xml:"GoodThreshold"`
	// The MaximumNonWorkHourResultsByDay element specifies the number of suggested results for meeting times outside regular working hours per day.
	MaximumNonWorkHourResultsByDay *MaximumNonWorkHourResultsByDay `xml:"MaximumNonWorkHourResultsByDay"`
	// The MaximumResultsByDay element specifies the number of suggested meeting times per a day returned in the response.
	MaximumResultsByDay *MaximumResultsByDay `xml:"MaximumResultsByDay"`
	// The MeetingDurationInMinutes element specifies the duration of the meeting to be suggested.
	MeetingDurationInMinutes *MeetingDurationInMinutes `xml:"MeetingDurationInMinutes"`
	// The MinimumSuggestionQuality element defines the quality of meeting suggestions to be returned in the response.
	MinimumSuggestionQuality *MinimumSuggestionQuality `xml:"MinimumSuggestionQuality"`
}

func (S *SuggestionsViewOptions) SetForMarshal() {
	S.XMLName.Local = "t:SuggestionsViewOptions"
}

func (S *SuggestionsViewOptions) GetSchema() *Schema {
	return &SchemaTypes
}
