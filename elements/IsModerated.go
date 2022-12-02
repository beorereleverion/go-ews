package elements

// The IsModerated element indicates whether the recipient's mailbox is being moderated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismoderated
type IsModerated bool

const (
	// false
	IsModeratedfalse IsModerated = false
	// true
	IsModeratedtrue IsModerated = true
)
