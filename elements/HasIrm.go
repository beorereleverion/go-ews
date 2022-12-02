package elements

// The HasIrm element specifies whether at least one message in the conversation and the current folder is an IRM protected message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hasirm
type HasIrm bool

const (
	// false
	HasIrmfalse HasIrm = false
	// true
	HasIrmtrue HasIrm = true
)
