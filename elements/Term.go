package elements

// The Term element specifies a highlighted term in a FindConversation or FindItem response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/term
type Term struct {
	// The Scope element specifies the string to be highlighted.
	Scope *ScopeHighlightTermType `xml:"t:Scope"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
