package elements

// The StopProcessingRules element indicates whether subsequent rules are to be evaluated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stopprocessingrules
type StopProcessingRules bool

const (
	// false
	StopProcessingRulesfalse StopProcessingRules = false
	// true
	StopProcessingRulestrue StopProcessingRules = true
)
