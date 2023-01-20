package elements

// The CalendarPermission element defines the access that a user has to a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarpermission
import "encoding/xml"

type CalendarPermission struct {
	XMLName xml.Name

	// The CalendarPermissionLevel element represents the permission level that a user has on a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarPermissionLevel *CalendarPermissionLevel `xml:"CalendarPermissionLevel"`
	// The CanCreateItems element indicates whether a user has permission to create items in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CanCreateItems *CanCreateItems `xml:"CanCreateItems"`
	// The CanCreateSubFolders element indicates whether a user has permission to create subfolders in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CanCreateSubFolders *CanCreateSubFolders `xml:"CanCreateSubFolders"`
	// The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DeleteItems *DeleteItems `xml:"DeleteItems"`
	// The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	EditItems *EditItems `xml:"EditItems"`
	// The IsFolderContact element indicates whether a user is a contact for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderContact *IsFolderContact `xml:"IsFolderContact"`
	// The IsFolderOwner element indicates whether a user is the owner of a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderOwner *IsFolderOwner `xml:"IsFolderOwner"`
	// The IsFolderVisible element indicates whether a user can view a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	IsFolderVisible *IsFolderVisible `xml:"IsFolderVisible"`
	// The ReadItems element indicates whether a user has permission to read items within a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	ReadItems *ReadItemsCalendarPermissionType `xml:"ReadItems"`
	// The UserId element identifies a delegate user or a user who has folder access permissions.
	UserId *UserId `xml:"UserId"`
}

func (C *CalendarPermission) SetForMarshal() {
	C.XMLName.Local = "t:CalendarPermission"
}

func (C *CalendarPermission) GetSchema() *Schema {
	return &SchemaTypes
}
