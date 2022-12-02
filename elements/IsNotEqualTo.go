package elements

// The IsNotEqualTo element represents a search expression that compares a property with either a constant value or another property and returns true if the values are not the same.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isnotequalto
type IsNotEqualTo struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"t:FieldURIOrConstant"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
