package elements

// The FolderChanges element represents a collection of changes for a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderchanges
type FolderChanges struct {
	// The FolderChange element represents a collection of changes to be performed on a single folder.
	FolderChange *FolderChange `xml:"t:FolderChange"`
}
