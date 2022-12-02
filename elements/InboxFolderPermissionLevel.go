package elements

// The InboxFolderPermissionLevel element contains the permissions for the default Inbox folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxfolderpermissionlevel
type InboxFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Inbox folder.
	InboxFolderPermissionLevelAuthor InboxFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Inbox folder.
	InboxFolderPermissionLevelCustom InboxFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Inbox folder.
	InboxFolderPermissionLevelEditor InboxFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Inbox folder.
	InboxFolderPermissionLevelNone InboxFolderPermissionLevel = `None`
	// The delegate user can read items in the Inbox folder.
	InboxFolderPermissionLevelReviewer InboxFolderPermissionLevel = `Reviewer`
)
