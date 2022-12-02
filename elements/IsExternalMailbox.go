package elements

// The IsExternalMailbox element indicates whether the mailbox is external to the organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isexternalmailbox
type IsExternalMailbox bool

const (
	// false
	IsExternalMailboxfalse IsExternalMailbox = false
	// true
	IsExternalMailboxtrue IsExternalMailbox = true
)
