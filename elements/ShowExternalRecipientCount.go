package elements

// The ShowExternalRecipientCount element indicates whether consumers of the GetMailTips operation have to show mail tips that indicate the number of external recipients to which a message is addressed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/showexternalrecipientcount
type ShowExternalRecipientCount bool

const (
	// false
	ShowExternalRecipientCountfalse ShowExternalRecipientCount = false
	// true
	ShowExternalRecipientCounttrue ShowExternalRecipientCount = true
)
