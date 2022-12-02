package elements

// The EmptyFolder element defines a request to empty a folder in a mailbox in the Exchange store. Optionally, subfolders can also be deleted when the folder is emptied.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emptyfolder
type EmptyFolder struct {
	// The FolderIds element contains an array of folder identifiers that are used to identify folders to copy, move, get, delete, or monitor for event notifications.
	FolderIds *FolderIds `xml:"m:FolderIds"`
	// Specifies whether subfolders are to be deleted. This attribute is required.
	DeleteSubFolders *string `xml:"DeleteSubFolders,attr"`
	// Specifies how a folder is emptied. This attribute is required.
	DeleteType *string `xml:"DeleteType,attr"`
}
