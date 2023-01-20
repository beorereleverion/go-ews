package elements

// The SuggestionsResponse element contains response status information and suggestion data for requested meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionsresponse
import "encoding/xml"

type SuggestionsResponse struct {
	XMLName xml.Name

	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"ResponseMessage"`
	// The SuggestionDayResultArray element contains an array of meeting suggestions organized by date.
	SuggestionDayResultArray *SuggestionDayResultArray `xml:"SuggestionDayResultArray"`
}

func (S *SuggestionsResponse) SetForMarshal() {
	S.XMLName.Local = "m:SuggestionsResponse"
}

func (S *SuggestionsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
