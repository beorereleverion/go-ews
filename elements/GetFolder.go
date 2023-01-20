package elements

// The GetFolder element defines a request to get a folder from a mailbox in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolder
import "encoding/xml"

type GetFolder struct {
	XMLName xml.Name

	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"FolderIds"`
	// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
	FolderShape *FolderShape `xml:"FolderShape"`
}

func (G *GetFolder) SetForMarshal() {
	G.XMLName.Local = "m:GetFolder"
}

func (G *GetFolder) GetSchema() *Schema {
	return &SchemaMessages
}
