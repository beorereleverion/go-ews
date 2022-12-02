package elements

// The FieldOrder element represents a single field by which to sort results and indicates the direction for the sort.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fieldorder
type FieldOrder struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
	// Describes the sort order direction. The following are the possible values:  - Ascending  - Descending
	Order *string `xml:"Order,attr"`
}
