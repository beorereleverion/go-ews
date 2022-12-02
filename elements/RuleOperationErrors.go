package elements

// The RuleOperationErrors element represents an array of rule validation errors on each rule field that has an error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ruleoperationerrors
type RuleOperationErrors struct {
	// The RuleOperationError element represents a rule operation error.
	RuleOperationError *RuleOperationError `xml:"m:RuleOperationError"`
}
