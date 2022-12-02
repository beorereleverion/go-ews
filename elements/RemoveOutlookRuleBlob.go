package elements

// The RemoveOutlookRuleBlob element indicates whether to remove the Microsoft Outlook rule blob.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeoutlookruleblob
type RemoveOutlookRuleBlob bool

const (
	// false
	RemoveOutlookRuleBlobfalse RemoveOutlookRuleBlob = false
	// true
	RemoveOutlookRuleBlobtrue RemoveOutlookRuleBlob = true
)
