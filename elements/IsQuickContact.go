package elements

// The IsQuickContact element specifies a Boolean value that indicates whether the underlying contact is a quick contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isquickcontact
type IsQuickContact bool

const (
	// false
	IsQuickContactfalse IsQuickContact = false
	// true
	IsQuickContacttrue IsQuickContact = true
)
