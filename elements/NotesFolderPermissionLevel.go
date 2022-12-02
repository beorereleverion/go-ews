package elements

// The NotesFolderPermissionLevel element contains the permissions for the default Notes folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notesfolderpermissionlevel
type NotesFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Notes folder.
	NotesFolderPermissionLevelAuthor NotesFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Notes folder.
	NotesFolderPermissionLevelCustom NotesFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Notes folder.
	NotesFolderPermissionLevelEditor NotesFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Notes folder.
	NotesFolderPermissionLevelNone NotesFolderPermissionLevel = `None`
	// The delegate user can read items in the Notes folder.
	NotesFolderPermissionLevelReviewer NotesFolderPermissionLevel = `Reviewer`
)
