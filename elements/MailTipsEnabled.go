package elements

// The MailTipsEnabled element indicates whether the mail tips service is available.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtipsenabled
type MailTipsEnabled bool

const (
	// false
	MailTipsEnabledfalse MailTipsEnabled = false
	// true
	MailTipsEnabledtrue MailTipsEnabled = true
)
