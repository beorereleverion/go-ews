package elements

// The DelegatePermissions element contains the delegate permission-level settings for a user. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegatepermissions
import "encoding/xml"

type DelegatePermissions struct {
	XMLName xml.Name

	// The CalendarFolderPermissionLevel element contains the permissions for the default Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarFolderPermissionLevel *CalendarFolderPermissionLevel `xml:"CalendarFolderPermissionLevel"`
	// The ContactsFolderPermissionLevel element contains the permissions for the default Contacts folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	ContactsFolderPermissionLevel *ContactsFolderPermissionLevel `xml:"ContactsFolderPermissionLevel"`
	// The InboxFolderPermissionLevel element contains the permissions for the default Inbox folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	InboxFolderPermissionLevel *InboxFolderPermissionLevel `xml:"InboxFolderPermissionLevel"`
	// The JournalFolderPermissionLevel element contains the permissions for the default Journal folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	JournalFolderPermissionLevel *JournalFolderPermissionLevel `xml:"JournalFolderPermissionLevel"`
	// The NotesFolderPermissionLevel element contains the permissions for the default Notes folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	NotesFolderPermissionLevel *NotesFolderPermissionLevel `xml:"NotesFolderPermissionLevel"`
	// The TasksFolderPermissionLevel element contains the permissions for the default Tasks folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	TasksFolderPermissionLevel *TasksFolderPermissionLevel `xml:"TasksFolderPermissionLevel"`
}

func (D *DelegatePermissions) SetForMarshal() {
	D.XMLName.Local = "t:DelegatePermissions"
}

func (D *DelegatePermissions) GetSchema() *Schema {
	return &SchemaTypes
}
