package elements

// The CreateContents element indicates whether a client can create a contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createcontents
type CreateContents bool

const (
	// false
	CreateContentsfalse CreateContents = false
	// true
	CreateContentstrue CreateContents = true
)
