package elements

// The Rules element contains an array of protection rules.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rules-ex15websvcsotherref
type Rules struct {
	// The Rule element contains a single protection rule.
	Rule *Rule `xml:"t:Rule"`
}
