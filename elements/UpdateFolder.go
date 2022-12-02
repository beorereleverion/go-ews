package elements

// The UpdateFolder element represents the operation that is used to update properties for a specified folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolder
type UpdateFolder struct {
	// The FolderChanges element represents a collection of changes for a folder.
	FolderChanges *FolderChanges `xml:"m:FolderChanges"`
}
