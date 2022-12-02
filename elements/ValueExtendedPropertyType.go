package elements

// The Value element specifies an array of extended properties for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-extendedpropertytype
type ValueExtendedPropertyType struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
	// The Values element contains a collection of values for an extended property.
	Values *Values `xml:"t:Values"`
}
