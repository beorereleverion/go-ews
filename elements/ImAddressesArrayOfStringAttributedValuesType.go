package elements

// The ImAddresses element specifies an array of instant message addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imaddresses-arrayofstringattributedvaluestype
type ImAddressesArrayOfStringAttributedValuesType struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
