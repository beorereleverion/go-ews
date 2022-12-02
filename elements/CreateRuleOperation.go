package elements

// The CreateRuleOperation element represents an operation to create a new Inbox rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createruleoperation
type CreateRuleOperation struct {
	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"t:Rule"`
}
