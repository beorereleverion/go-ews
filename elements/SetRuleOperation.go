package elements

// The SetRuleOperation element represents an operation to update an existing rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setruleoperation
type SetRuleOperation struct {
	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"t:Rule"`
}
