package elements

// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduriorconstant
type FieldURIOrConstant struct {
	// The Constant element identifies a constant value in a restriction.
	Constant *Constant `xml:"t:Constant"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
