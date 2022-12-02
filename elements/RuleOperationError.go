package elements

// The RuleOperationError element represents a rule operation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ruleoperationerror
type RuleOperationError struct {
	// The OperationIndex element specifies the index of the operation in the request that caused the rule operation error.
	OperationIndex *OperationIndex `xml:"m:OperationIndex"`
	// The ValidationErrors element represents an array of rule validation errors on each rule field that has an error.
	ValidationErrors *ValidationErrors `xml:"m:ValidationErrors"`
}
