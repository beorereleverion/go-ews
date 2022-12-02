package elements

// The And element specifies that all child elements must match to evaluate to true.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/and-protectionruleandtype
type AndProtectionRuleAndType struct {
	// The AllInternal element evaluates to true if all recipients of an e-mail message are internal to the sender's organization.
	AllInternal *AllInternal `xml:"t:AllInternal"`
	// The RecipientIs element specifies that any recipient of the e-mail message matches any of the specified recipients in the child Value (ProtectionRuleValueType) elements.
	RecipientIs *RecipientIs `xml:"t:RecipientIs"`
	// The SenderDepartments element specifies that the department of the sender matches any of the specified departments in the child Value (ProtectionRuleValueType) elements.
	SenderDepartments *SenderDepartments `xml:"t:SenderDepartments"`
	// The True element specifies a condition that always matches.
	True *True `xml:"t:True"`
	// The Value element identifies a single recipient or sender department.
	Value *ValueProtectionRuleValueType `xml:"t:Value"`
}
