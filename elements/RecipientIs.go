package elements

// The RecipientIs element specifies that any recipient of the e-mail message matches any of the specified recipients in the child Value (ProtectionRuleValueType) elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipientis
type RecipientIs struct {
	// The Value element identifies a single recipient or sender department.
	Value *ValueProtectionRuleValueType `xml:"t:Value"`
}
