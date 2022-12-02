package elements

// The Delete element identifies a single folder to delete in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delete-foldersync
type DeleteFolderSync struct {
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
}
