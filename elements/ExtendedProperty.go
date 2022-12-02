package elements

// The ExtendedProperty element identifies extended MAPI properties on folders and items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedproperty
type ExtendedProperty struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
	// The Values element contains a collection of values for an extended property.
	Values *Values `xml:"t:Values"`
}
