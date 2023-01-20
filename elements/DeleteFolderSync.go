package elements

// The Delete element identifies a single folder to delete in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delete-foldersync
import "encoding/xml"

type DeleteFolderSync struct {
	XMLName xml.Name

	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (D *DeleteFolderSync) SetForMarshal() {
	D.XMLName.Local = "t:Delete"
}

func (D *DeleteFolderSync) GetSchema() *Schema {
	return &SchemaTypes
}
