package elements

// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additionalproperties
type AdditionalProperties struct {
	// Identifies extended MAPI properties to get, set, or create.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// Identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// Identifies frequently referenced dictionary properties by URI.
	IndexedFieldURI *string `xml:"t:IndexedFieldURI"`
}
