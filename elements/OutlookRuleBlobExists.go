package elements

// The OutlookRuleBlobExists element indicates whether a Microsoft Outlook rule blob exists in the user's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/outlookruleblobexists
type OutlookRuleBlobExists bool

const (
	// false
	OutlookRuleBlobExistsfalse OutlookRuleBlobExists = false
	// true
	OutlookRuleBlobExiststrue OutlookRuleBlobExists = true
)
