package elements

// The JournalFolderPermissionLevel element contains the permissions for the default Journal folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/journalfolderpermissionlevel
type JournalFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Journal folder.
	JournalFolderPermissionLevelAuthor JournalFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Journal folder.
	JournalFolderPermissionLevelCustom JournalFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Journal folder.
	JournalFolderPermissionLevelEditor JournalFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Journal folder.
	JournalFolderPermissionLevelNone JournalFolderPermissionLevel = `None`
	// The delegate user can read items in the Journal folder.
	JournalFolderPermissionLevelReviewer JournalFolderPermissionLevel = `Reviewer`
)
