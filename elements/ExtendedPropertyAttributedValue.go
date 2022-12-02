package elements

// The ExtendedPropertyAttributedValue element specifies extended properties for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedpropertyattributedvalue
type ExtendedPropertyAttributedValue struct {
	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"t:Attributions"`
	// The Value element specifies an array of extended properties for a persona.
	Value *ValueExtendedPropertyType `xml:"t:Value"`
}
