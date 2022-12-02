package elements

// The IsMembershipGroup element specifies a Boolean value that indicates whether the entity is a distribution group or a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismembershipgroup
type IsMembershipGroup bool

const (
	// false
	IsMembershipGroupfalse IsMembershipGroup = false
	// true
	IsMembershipGrouptrue IsMembershipGroup = true
)
