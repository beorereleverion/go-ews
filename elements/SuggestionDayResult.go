package elements

// The SuggestionDayResult element represents a single day that contains suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestiondayresult
import "encoding/xml"

type SuggestionDayResult struct {
	XMLName xml.Name

	// The Date element represents the date that contains the suggested meeting times.
	Date *Date `xml:"Date"`
	// The DayQuality element represents the quality of the day for containing quality suggested meeting times.
	DayQuality *DayQuality `xml:"DayQuality"`
	// The SuggestionArray element contains an array of meeting suggestions.
	SuggestionArray *SuggestionArray `xml:"SuggestionArray"`
}

func (S *SuggestionDayResult) SetForMarshal() {
	S.XMLName.Local = "t:SuggestionDayResult"
}

func (S *SuggestionDayResult) GetSchema() *Schema {
	return &SchemaTypes
}
