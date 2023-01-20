package elements

// The Rule element contains a single protection rule.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rule
import "encoding/xml"

type Rule struct {
	XMLName xml.Name

	// The Action element identifies what action must be executed if the condition part of the rule matches.
	Action *ActionProtectionRuleActionType `xml:"Action"`
	// The Condition element identifies the condition that must be satisfied for the action part of the rule to be executed.
	Condition *Condition `xml:"Condition"`
	// Identifies the name of the rule. A required attribute of type string with a minimum length of 1.
	Name *string `xml:"Name,attr"`
	// Specifies the rule priority. A required attribute of type int with a minimum value of 1.
	Priority *string `xml:"Priority,attr"`
	// Specifies whether the rule is mandatory. If the rule is mandatory, this attribute value must be false. A required attribute of type Boolean.
	UserOverridable *string `xml:"UserOverridable,attr"`
}

func (R *Rule) SetForMarshal() {
	R.XMLName.Local = "t:Rule"
}

func (R *Rule) GetSchema() *Schema {
	return &SchemaTypes
}
