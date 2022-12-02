package elements

// The FolderNames element contains an array of named managed folders to add to a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/foldernames
type FolderNames struct {
	// The FolderName element identifies a single managed custom folder to add to a mailbox.
	FolderName *FolderName `xml:"t:FolderName"`
}
