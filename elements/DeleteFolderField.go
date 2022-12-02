package elements

// The DeleteFolderField element represents an operation to delete a given property from a folder during an UpdateFolder call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolderfield
type DeleteFolderField struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
}
