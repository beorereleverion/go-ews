package elements

// The DetailedSuggestionsWindow element identifies the time span that is queried for detailed information about suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/detailedsuggestionswindow
import "encoding/xml"

type DetailedSuggestionsWindow struct {
	XMLName xml.Name

	// The EndTime element represents the end of a time span.
	EndTime *EndTime `xml:"EndTime"`
	// The StartTime element represents the start of a time span.
	StartTime *StartTime `xml:"StartTime"`
}

func (D *DetailedSuggestionsWindow) SetForMarshal() {
	D.XMLName.Local = "t:DetailedSuggestionsWindow"
}

func (D *DetailedSuggestionsWindow) GetSchema() *Schema {
	return &SchemaTypes
}
