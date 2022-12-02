package elements

// The PermanentDelete element indicates whether messages are to be permanently deleted and not saved to the Deleted Items folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/permanentdelete
type PermanentDelete bool

const (
	// false
	PermanentDeletefalse PermanentDelete = false
	// true
	PermanentDeletetrue PermanentDelete = true
)
