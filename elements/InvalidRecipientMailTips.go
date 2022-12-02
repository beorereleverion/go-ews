package elements

// The InvalidRecipient element indicates whether the recipient is invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/invalidrecipient-mailtips
type InvalidRecipientMailTips bool

const (
	// false
	InvalidRecipientMailTipsfalse InvalidRecipientMailTips = false
	// true
	InvalidRecipientMailTipstrue InvalidRecipientMailTips = true
)
