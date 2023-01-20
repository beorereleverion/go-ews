package elements

// The SyncFolderHierarchy element defines a request to synchronize a folder hierarchy on a client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderhierarchy
import "encoding/xml"

type SyncFolderHierarchy struct {
	XMLName xml.Name

	// The FolderShape element identifies the folder properties to include in a GetFolder, FindFolder, or SyncFolderHierarchy response.
	FolderShape *FolderShape `xml:"FolderShape"`
	// The SyncFolderHierarchy element defines a request to synchronize a folder hierarchy on a client.
	SyncFolderHierarchy *SyncFolderHierarchy `xml:"SyncFolderHierarchy"`
	// The SyncFolderId element represents the folder that contains the items to synchronize.
	SyncFolderId *SyncFolderId `xml:"SyncFolderId"`
	// The SyncState element contains a base64-encoded form of the synchronization data that is updated after each successful request. This is used to identify the synchronization state.
	SyncState *SyncState `xml:"SyncState"`
}

func (S *SyncFolderHierarchy) SetForMarshal() {
	S.XMLName.Local = "m:SyncFolderHierarchy"
}

func (S *SyncFolderHierarchy) GetSchema() *Schema {
	return &SchemaMessages
}
