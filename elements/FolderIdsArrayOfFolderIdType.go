package elements

// The FolderIds element contains a list of folder identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderids-arrayoffolderidtype
type FolderIdsArrayOfFolderIdType struct {
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
}
