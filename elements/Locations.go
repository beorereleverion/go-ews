package elements

// The Locations element specifies an array of location values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/locations
type Locations struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
