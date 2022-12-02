package elements

// The DeleteItemField element represents an operation to delete a given property from an item during an UpdateItem call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemfield
type DeleteItemField struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
