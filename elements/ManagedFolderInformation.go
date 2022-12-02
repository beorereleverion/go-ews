package elements

// The ManagedFolderInformation element contains information about a managed custom folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managedfolderinformation
type ManagedFolderInformation struct {
	// The CanDelete element indicates whether a managed folder can be deleted by a customer.
	CanDelete *CanDelete `xml:"t:CanDelete"`
	// The CanRenameOrMove element indicates whether a managed folder can be renamed or moved by the customer.
	CanRenameOrMove *CanRenameOrMove `xml:"t:CanRenameOrMove"`
	// The Comment element contains the comment that is associated with a managed folder.
	Comment *Comment `xml:"t:Comment"`
	// The FolderSize element describes the total size of all the contents of a managed folder.
	FolderSize *FolderSize `xml:"t:FolderSize"`
	// The HasQuota element indicates whether the managed folder has a quota.
	HasQuota *HasQuota `xml:"t:HasQuota"`
	// The HomePage element specifies the URL that will be the default home page for the managed folder.
	HomePage *HomePage `xml:"t:HomePage"`
	// The IsManagedFoldersRoot element indicates whether the managed folder is the root for all managed folders.
	IsManagedFoldersRoot *IsManagedFoldersRoot `xml:"t:IsManagedFoldersRoot"`
	// The ManagedFolderId element contains the folder ID of the managed folder.
	ManagedFolderId *ManagedFolderId `xml:"t:ManagedFolderId"`
	// The MustDisplayComment element indicates whether the managed folder comment must be displayed.
	MustDisplayComment *MustDisplayComment `xml:"t:MustDisplayComment"`
	// The StorageQuota element describes the storage quota for the managed folder.
	StorageQuota *StorageQuota `xml:"t:StorageQuota"`
}
