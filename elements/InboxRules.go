package elements

// The InboxRules element represents an array of rules in the user's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxrules
type InboxRules struct {
	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"t:Rule"`
}
