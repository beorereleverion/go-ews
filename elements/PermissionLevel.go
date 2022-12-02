package elements

// The PermissionLevel element represents the permission level that a user has on a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permissionlevel
type PermissionLevel string

const (
	// Author
	PermissionLevelAuthor PermissionLevel = `Author`
	// Contributor
	PermissionLevelContributor PermissionLevel = `Contributor`
	// Custom
	PermissionLevelCustom PermissionLevel = `Custom`
	// Editor
	PermissionLevelEditor PermissionLevel = `Editor`
	// None
	PermissionLevelNone PermissionLevel = `None`
	// NoneditingAuthor
	PermissionLevelNoneditingAuthor PermissionLevel = `NoneditingAuthor`
	// Owner
	PermissionLevelOwner PermissionLevel = `Owner`
	// PublishingAuthor
	PermissionLevelPublishingAuthor PermissionLevel = `PublishingAuthor`
	// PublishingEditor
	PermissionLevelPublishingEditor PermissionLevel = `PublishingEditor`
	// Reviewer
	PermissionLevelReviewer PermissionLevel = `Reviewer`
)
