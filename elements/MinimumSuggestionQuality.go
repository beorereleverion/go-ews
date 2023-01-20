package elements

// The MinimumSuggestionQuality element defines the quality of meeting suggestions to be returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/minimumsuggestionquality
import "encoding/xml"

type MinimumSuggestionQuality struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// 0% of the attendees have a conflict with the suggested meeting time.
	MinimumSuggestionQualityExcellent string = `Excellent`
	// The percentage that is considered fair is set by using the GoodThreshold element.
	MinimumSuggestionQualityFair string = `Fair`
	// The percentage that is considered good is set by using the GoodThreshold element.
	MinimumSuggestionQualityGood string = `Good`
	// 50% or more of the attendees have a conflict with the suggested meeting time.
	MinimumSuggestionQualityPoor string = `Poor`
)

func (M *MinimumSuggestionQuality) SetForMarshal() {
	M.XMLName.Local = "t:MinimumSuggestionQuality"
}

func (M *MinimumSuggestionQuality) GetSchema() *Schema {
	return &SchemaTypes
}
