package elements

// The Condition element identifies the condition that must be satisfied for the action part of the rule to be executed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/condition
import "encoding/xml"

type Condition struct {
	XMLName xml.Name

	// The AllInternal element evaluates to true if all recipients of an e-mail message are internal to the sender's organization.
	AllInternal *AllInternal `xml:"AllInternal"`
	// The And element specifies that all child elements must match to evaluate to true.
	And *AndProtectionRuleAndType `xml:"And"`
	// The RecipientIs element specifies that any recipient of the e-mail message matches any of the specified recipients in the child Value (ProtectionRuleValueType) elements.
	RecipientIs *RecipientIs `xml:"RecipientIs"`
	// The SenderDepartments element specifies that the department of the sender matches any of the specified departments in the child Value (ProtectionRuleValueType) elements.
	SenderDepartments *SenderDepartments `xml:"SenderDepartments"`
	// The True element specifies a condition that always matches.
	True *True `xml:"True"`
	// The Value element identifies a single recipient or sender department.
	Value *ValueProtectionRuleValueType `xml:"Value"`
}

func (C *Condition) SetForMarshal() {
	C.XMLName.Local = "t:Condition"
}

func (C *Condition) GetSchema() *Schema {
	return &SchemaTypes
}
