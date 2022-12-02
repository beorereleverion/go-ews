package elements

// The Permission element defines the access that a user has to a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permission
type Permission struct {
	// The CanCreateItems element indicates whether a user has permission to create items in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CanCreateItems *CanCreateItems `xml:"t:CanCreateItems"`
	// The CanCreateSubFolders element indicates whether a user has permission to create subfolders in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CanCreateSubFolders *CanCreateSubFolders `xml:"t:CanCreateSubFolders"`
	// The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DeleteItems *DeleteItems `xml:"t:DeleteItems"`
	// The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	EditItems *EditItems `xml:"t:EditItems"`
	// The IsFolderContact element indicates whether a user is a contact for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderContact *IsFolderContact `xml:"t:IsFolderContact"`
	// The IsFolderOwner element indicates whether a user is the owner of a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderOwner *IsFolderOwner `xml:"t:IsFolderOwner"`
	// The IsFolderVisible element indicates whether a user can view a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderVisible *IsFolderVisible `xml:"t:IsFolderVisible"`
	// The PermissionLevel element represents the permission level that a user has on a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	PermissionLevel *PermissionLevel `xml:"t:PermissionLevel"`
	// The ReadItems element indicates whether a user has permission to read items within a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	ReadItems *ReadItemsPermissionType `xml:"t:ReadItems"`
	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"t:UserId"`
}
