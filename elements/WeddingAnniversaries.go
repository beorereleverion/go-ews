package elements

// The WeddingAnniversaries element specifies an array of wedding anniversary dates, stored as strings, and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weddinganniversaries
type WeddingAnniversaries struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
