package elements

// The TasksFolderPermissionLevel element contains the permissions for the default Tasks folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasksfolderpermissionlevel
type TasksFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Tasks folder.
	TasksFolderPermissionLevelAuthor TasksFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Tasks folder.
	TasksFolderPermissionLevelCustom TasksFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Tasks folder.
	TasksFolderPermissionLevelEditor TasksFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Tasks folder.
	TasksFolderPermissionLevelNone TasksFolderPermissionLevel = `None`
	// The delegate user can read items in the Tasks folder.
	TasksFolderPermissionLevelReviewer TasksFolderPermissionLevel = `Reviewer`
)
