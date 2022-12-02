package elements

// The HiddenRecipient element indicates that the recipient was added by an organization policy that should be hidden from unprivileged users.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/hiddenrecipient
type HiddenRecipient bool

const (
	// false
	HiddenRecipientfalse HiddenRecipient = false
	// true
	HiddenRecipienttrue HiddenRecipient = true
)
