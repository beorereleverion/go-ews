package elements

// The SuggestionArray element contains an array of meeting suggestions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/suggestionarray
import "encoding/xml"

type SuggestionArray struct {
	XMLName xml.Name

	// The Suggestion element represents a single meeting suggestion.
	Suggestion *Suggestion `xml:"Suggestion"`
}

func (S *SuggestionArray) SetForMarshal() {
	S.XMLName.Local = "t:SuggestionArray"
}

func (S *SuggestionArray) GetSchema() *Schema {
	return &SchemaTypes
}
