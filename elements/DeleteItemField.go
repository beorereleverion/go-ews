package elements

// The DeleteItemField element represents an operation to delete a given property from an item during an UpdateItem call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemfield
import "encoding/xml"

type DeleteItemField struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}
