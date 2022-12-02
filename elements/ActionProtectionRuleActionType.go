package elements

// The Action element identifies what action must be executed if the condition part of the rule matches.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/action-protectionruleactiontype
type ActionProtectionRuleActionType struct {
	// The Argument element specifies arguments to the action.
	Argument *Argument `xml:"t:Argument"`
	// Identifies the name of the action.
	Name *string `xml:"Name,attr"`
}
