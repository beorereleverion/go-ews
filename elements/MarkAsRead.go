package elements

// The MarkAsRead element indicates whether messages are to be marked as read.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markasread
type MarkAsRead bool

const (
	// false
	MarkAsReadfalse MarkAsRead = false
	// true
	MarkAsReadtrue MarkAsRead = true
)
