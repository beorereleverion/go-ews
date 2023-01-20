package elements

// The DeleteRuleOperation element contains an operation to delete an existing Inbox rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteruleoperation
import "encoding/xml"

type DeleteRuleOperation struct {
	XMLName xml.Name

	// The RuleId element specifies a rule identifier.
	RuleId *RuleId `xml:"RuleId"`
}

func (D *DeleteRuleOperation) SetForMarshal() {
	D.XMLName.Local = "t:DeleteRuleOperation"
}

func (D *DeleteRuleOperation) GetSchema() *Schema {
	return &SchemaTypes
}
