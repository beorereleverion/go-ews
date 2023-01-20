package elements

// The MoveFolder element defines a request to move a folder in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/movefolder
import "encoding/xml"

type MoveFolder struct {
	XMLName xml.Name

	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"FolderIds"`
	// The ToFolderId element represents the destination folder for a copied or moved item or folder.
	ToFolderId *ToFolderId `xml:"ToFolderId"`
}

func (M *MoveFolder) SetForMarshal() {
	M.XMLName.Local = "m:MoveFolder"
}

func (M *MoveFolder) GetSchema() *Schema {
	return &SchemaMessages
}
