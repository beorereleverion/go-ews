package elements

// The Delete element indicates whether a client can delete a folder or item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delete
type Delete bool

const (
	// false
	Deletefalse Delete = false
	// true
	Deletetrue Delete = true
)
