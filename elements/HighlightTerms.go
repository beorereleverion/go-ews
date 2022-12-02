package elements

// The HighlightTerms element identifies the highlighted terms returned in a FindItem operation and a FindConversation operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/highlightterms
type HighlightTerms struct {
	// The Term element specifies a highlighted term in a FindConversation or FindItem response.
	Term *Term `xml:"t:Term"`
}
