package elements

// The HighlightTerms element identifies the highlighted terms returned in a FindItem operation and a FindConversation operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/highlightterms
import "encoding/xml"

type HighlightTerms struct {
	XMLName xml.Name

	// The Term element specifies a highlighted term in a FindConversation or FindItem response.
	Term *Term `xml:"Term"`
}

func (H *HighlightTerms) SetForMarshal() {
	H.XMLName.Local = "m:HighlightTerms"
}

func (H *HighlightTerms) GetSchema() *Schema {
	return &SchemaMessages
}
