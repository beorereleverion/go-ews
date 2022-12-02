package elements

// The DeleteRuleOperation element contains an operation to delete an existing Inbox rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteruleoperation
type DeleteRuleOperation struct {
	// The RuleId element specifies a rule identifier.
	RuleId *RuleId `xml:"m:RuleId"`
}
