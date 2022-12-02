package elements

// The GlobalHasIrm element specifies whether at least one message in the conversation and across all folders is an IRM protected message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalhasirm
type GlobalHasIrm bool

const (
	// false
	GlobalHasIrmfalse GlobalHasIrm = false
	// true
	GlobalHasIrmtrue GlobalHasIrm = true
)
