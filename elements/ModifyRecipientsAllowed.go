package elements

// The ModifyRecipientsAllowed element specifies whether modification of the recipients is enabled.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modifyrecipientsallowed
type ModifyRecipientsAllowed bool

const (
	// false
	ModifyRecipientsAllowedfalse ModifyRecipientsAllowed = false
	// true
	ModifyRecipientsAllowedtrue ModifyRecipientsAllowed = true
)
