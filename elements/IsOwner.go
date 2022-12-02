package elements

// The IsOwner element specifies whether the specified email user is the owner.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isowner
type IsOwner bool

const (
	// false
	IsOwnerfalse IsOwner = false
	// true
	IsOwnertrue IsOwner = true
)
