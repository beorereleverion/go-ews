package elements

// The DistinguishedUser element identifies Anonymous and Default user accounts for delegate access. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguisheduser
type DistinguishedUser string

const (
	// anonymous
	DistinguishedUseranonymous DistinguishedUser = `anonymous`
	// default
	DistinguishedUserdefault DistinguishedUser = `default`
)
