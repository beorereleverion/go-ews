package elements

// The StringArrayAttributedValue element specifies an instance of an array of string data for a persona element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stringarrayattributedvalue
type StringArrayAttributedValue struct {
	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"t:Attributions"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
