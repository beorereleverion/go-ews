package elements

// The Excludes element performs a bitwise mask of the specified property and a supplied value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/excludes
type Excludes struct {
	// The Bitmask element represents a hexadecimal or decimal mask to be used during an Excludes restriction operation.
	Bitmask *Bitmask `xml:"t:Bitmask"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"t:Excludes"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
