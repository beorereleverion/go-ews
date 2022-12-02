package elements

// The CalendarPermissionLevel element represents the permission level that a user has on a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/calendarpermissionlevel
type CalendarPermissionLevel string

const (
	// Author
	CalendarPermissionLevelAuthor CalendarPermissionLevel = `Author`
	// Contributor
	CalendarPermissionLevelContributor CalendarPermissionLevel = `Contributor`
	// Custom
	CalendarPermissionLevelCustom CalendarPermissionLevel = `Custom`
	// Editor
	CalendarPermissionLevelEditor CalendarPermissionLevel = `Editor`
	// FreeBusyTimeAndSubjectAndLocation
	CalendarPermissionLevelFreeBusyTimeAndSubjectAndLocation CalendarPermissionLevel = `FreeBusyTimeAndSubjectAndLocation`
	// FreeBusyTimeOnly
	CalendarPermissionLevelFreeBusyTimeOnly CalendarPermissionLevel = `FreeBusyTimeOnly`
	// None
	CalendarPermissionLevelNone CalendarPermissionLevel = `None`
	// NoneditingAuthor
	CalendarPermissionLevelNoneditingAuthor CalendarPermissionLevel = `NoneditingAuthor`
	// Owner
	CalendarPermissionLevelOwner CalendarPermissionLevel = `Owner`
	// PublishingAuthor
	CalendarPermissionLevelPublishingAuthor CalendarPermissionLevel = `PublishingAuthor`
	// PublishingEditor
	CalendarPermissionLevelPublishingEditor CalendarPermissionLevel = `PublishingEditor`
	// Reviewer
	CalendarPermissionLevelReviewer CalendarPermissionLevel = `Reviewer`
)
