package elements

// The Rule element contains a single rule and represents a rule in a user's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rule-ruletype
type RuleRuleType struct {
	// The Actions element represents the set of actions that are available to be taken on a message when the conditions are fulfilled.
	Actions *Actions `xml:"t:Actions"`
	// The Conditions element identifies the conditions that, when fulfilled, will trigger the rule actions for a rule.
	Conditions *Conditions `xml:"t:Conditions"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"t:DisplayName"`
	// The Exceptions element identifies the exceptions that represent all the available rule exception conditions for an Inbox rule.
	Exceptions *Exceptions `xml:"t:Exceptions"`
	// The IsEnabled element indicates whether the rule is enabled.
	IsEnabled *IsEnabled `xml:"m:IsEnabled"`
	// The IsInError element indicates whether the rule is in an error condition.
	IsInError *IsInError `xml:"m:IsInError"`
	// The IsNotSupported element indicates whether the rule cannot be modified by using the managed code APIs.
	IsNotSupported *IsNotSupported `xml:"m:IsNotSupported"`
	// The Priority element indicates the order in which a rule is to be run.
	Priority *Priority `xml:"m:Priority"`
	// The RuleId element specifies a rule identifier.
	RuleId *RuleId `xml:"m:RuleId"`
}
