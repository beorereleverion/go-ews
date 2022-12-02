package elements

// The ContactsFolderPermissionLevel element contains the permissions for the default Contacts folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactsfolderpermissionlevel
type ContactsFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Contacts folder.
	ContactsFolderPermissionLevelAuthor ContactsFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Contacts folder.
	ContactsFolderPermissionLevelCustom ContactsFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Contacts folder.
	ContactsFolderPermissionLevelEditor ContactsFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Contacts folder.
	ContactsFolderPermissionLevelNone ContactsFolderPermissionLevel = `None`
	// The delegate user can read items in the Contacts folder.
	ContactsFolderPermissionLevelReviewer ContactsFolderPermissionLevel = `Reviewer`
)
