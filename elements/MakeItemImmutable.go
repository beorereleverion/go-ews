package elements

// The MakeItemImmutable element specifies a Boolean value that indicates whether an item should be made read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/makeitemimmutable
type MakeItemImmutable bool

const (
	// false
	MakeItemImmutablefalse MakeItemImmutable = false
	// true
	MakeItemImmutabletrue MakeItemImmutable = true
)
