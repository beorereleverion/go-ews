package elements

// The SetRuleOperation element represents an operation to update an existing rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setruleoperation
import "encoding/xml"

type SetRuleOperation struct {
	XMLName xml.Name

	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"Rule"`
}

func (S *SetRuleOperation) SetForMarshal() {
	S.XMLName.Local = "t:SetRuleOperation"
}

func (S *SetRuleOperation) GetSchema() *Schema {
	return &SchemaTypes
}
