package elements

// The CreateRuleOperation element represents an operation to create a new Inbox rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createruleoperation
import "encoding/xml"

type CreateRuleOperation struct {
	XMLName xml.Name

	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"Rule"`
}

func (C *CreateRuleOperation) SetForMarshal() {
	C.XMLName.Local = "t:CreateRuleOperation"
}

func (C *CreateRuleOperation) GetSchema() *Schema {
	return &SchemaTypes
}
