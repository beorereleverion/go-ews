package elements

// The ImAddresses2 element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddresses2
type ImAddresses2 struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
