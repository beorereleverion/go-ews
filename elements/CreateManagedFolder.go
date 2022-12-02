package elements

// The CreateManagedFolder element defines a request to add managed custom folders to a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createmanagedfolder
type CreateManagedFolder struct {
	// The FolderNames element contains an array of named managed folders to add to a mailbox.
	FolderNames *FolderNames `xml:"m:FolderNames"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"t:Mailbox"`
}
