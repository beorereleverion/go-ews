package elements

// The Updates element contains a set of elements that define append, set, and delete changes to folder properties.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updates-folder
import "encoding/xml"

type UpdatesFolder struct {
	XMLName xml.Name

	// The AppendToFolderField element is not implemented. Any request that uses this element will always return an error response.
	AppendToFolderField *AppendToFolderField `xml:"AppendToFolderField"`
	// The DeleteFolderField element represents an operation to delete a given property from a folder during an UpdateFolder call.
	DeleteFolderField *DeleteFolderField `xml:"DeleteFolderField"`
	// The SetFolderField element represents an update that sets the value for a single property on a folder in an UpdateFolder operation.
	SetFolderField *SetFolderField `xml:"SetFolderField"`
}

func (U *UpdatesFolder) SetForMarshal() {
	U.XMLName.Local = "t:Updates"
}

func (U *UpdatesFolder) GetSchema() *Schema {
	return &SchemaTypes
}
