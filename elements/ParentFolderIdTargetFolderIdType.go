package elements

// The ParentFolderId element identifies the folder in which a new folder is created or the folder to search for the FindConversation operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/parentfolderid-targetfolderidtype
type ParentFolderIdTargetFolderIdType struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
}
