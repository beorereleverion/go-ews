package elements

// The ManagedFolderInformation element contains information about a managed custom folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managedfolderinformation
import "encoding/xml"

type ManagedFolderInformation struct {
	XMLName xml.Name

	// The CanDelete element indicates whether a managed folder can be deleted by a customer.
	CanDelete *CanDelete `xml:"CanDelete"`
	// The CanRenameOrMove element indicates whether a managed folder can be renamed or moved by the customer.
	CanRenameOrMove *CanRenameOrMove `xml:"CanRenameOrMove"`
	// The Comment element contains the comment that is associated with a managed folder.
	Comment *Comment `xml:"Comment"`
	// The FolderSize element describes the total size of all the contents of a managed folder.
	FolderSize *FolderSize `xml:"FolderSize"`
	// The HasQuota element indicates whether the managed folder has a quota.
	HasQuota *HasQuota `xml:"HasQuota"`
	// The HomePage element specifies the URL that will be the default home page for the managed folder.
	HomePage *HomePage `xml:"HomePage"`
	// The IsManagedFoldersRoot element indicates whether the managed folder is the root for all managed folders.
	IsManagedFoldersRoot *IsManagedFoldersRoot `xml:"IsManagedFoldersRoot"`
	// The ManagedFolderId element contains the folder ID of the managed folder.
	ManagedFolderId *ManagedFolderId `xml:"ManagedFolderId"`
	// The MustDisplayComment element indicates whether the managed folder comment must be displayed.
	MustDisplayComment *MustDisplayComment `xml:"MustDisplayComment"`
	// The StorageQuota element describes the storage quota for the managed folder.
	StorageQuota *StorageQuota `xml:"StorageQuota"`
}

func (M *ManagedFolderInformation) SetForMarshal() {
	M.XMLName.Local = "t:ManagedFolderInformation"
}

func (M *ManagedFolderInformation) GetSchema() *Schema {
	return &SchemaTypes
}
