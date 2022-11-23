package elements

// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderids
type FolderIds struct {
	// Contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId"`
	// Identifies Microsoft Exchange Server folders that can be referenced by name.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}
