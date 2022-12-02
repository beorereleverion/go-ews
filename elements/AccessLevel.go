package elements

// The AccessLevel element specifies the access level for an online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/accesslevel
type AccessLevel string

const (
	// The access level is open to all.
	AccessLevelEveryone AccessLevel = `Everyone`
	// The access level is internal only.
	AccessLevelInternal AccessLevel = `Internal`
	// The access level is invited participants only.
	AccessLevelInvited AccessLevel = `Invited`
	// The access level is locked.
	AccessLevelLocked AccessLevel = `Locked`
)
