package elements

// The DeleteFolderField element represents an operation to delete a given property from a folder during an UpdateFolder call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolderfield
import "encoding/xml"

type DeleteFolderField struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (D *DeleteFolderField) SetForMarshal() {
	D.XMLName.Local = "t:DeleteFolderField"
}

func (D *DeleteFolderField) GetSchema() *Schema {
	return &SchemaTypes
}
