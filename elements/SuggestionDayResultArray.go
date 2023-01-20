package elements

// The SuggestionDayResultArray element contains an array of meeting suggestions organized by date.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestiondayresultarray
import "encoding/xml"

type SuggestionDayResultArray struct {
	XMLName xml.Name

	// The SuggestionDayResult element represents a single day that contains suggested meeting times.
	SuggestionDayResult *SuggestionDayResult `xml:"SuggestionDayResult"`
}

func (S *SuggestionDayResultArray) SetForMarshal() {
	S.XMLName.Local = "m:SuggestionDayResultArray"
}

func (S *SuggestionDayResultArray) GetSchema() *Schema {
	return &SchemaMessages
}
