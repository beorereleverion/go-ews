package elements

// The CreateAssociated element indicates whether a client can create an associated contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createassociated
type CreateAssociated bool

const (
	// false
	CreateAssociatedfalse CreateAssociated = false
	// true
	CreateAssociatedtrue CreateAssociated = true
)
