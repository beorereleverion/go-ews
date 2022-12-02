package elements

// The Read element indicates whether a client can read a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/read
type Read bool

const (
	// false
	Readfalse Read = false
	// true
	Readtrue Read = true
)
