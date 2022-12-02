package elements

// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stringattributedvalue
type StringAttributedValue struct {
	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"t:Attributions"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
