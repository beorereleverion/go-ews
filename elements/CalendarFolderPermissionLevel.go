package elements

// The CalendarFolderPermissionLevel element contains the permissions for the default Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarfolderpermissionlevel
type CalendarFolderPermissionLevel string

const (
	// The delegate user can read and create items in the Calendar folder.
	CalendarFolderPermissionLevelAuthor CalendarFolderPermissionLevel = `Author`
	// The delegate user has custom access permissions to the Calendar folder.
	CalendarFolderPermissionLevelCustom CalendarFolderPermissionLevel = `Custom`
	// The delegate user can read, create, and modify items in the Calendar folder.
	CalendarFolderPermissionLevelEditor CalendarFolderPermissionLevel = `Editor`
	// The delegate user has no access permissions to the Calendar folder.
	CalendarFolderPermissionLevelNone CalendarFolderPermissionLevel = `None`
	// The delegate user can read items in the Calendar folder.
	CalendarFolderPermissionLevelReviewer CalendarFolderPermissionLevel = `Reviewer`
)
