package elements

// The ExpandGroupMembership element indicates whether to expand the membership of the group returned from a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expandgroupmembership
type ExpandGroupMembership bool

const (
	// false
	ExpandGroupMembershipfalse ExpandGroupMembership = false
	// true
	ExpandGroupMembershiptrue ExpandGroupMembership = true
)
