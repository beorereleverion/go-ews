package elements

// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additionalproperties
type AdditionalProperties struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
