package elements

// The Action element identifies what action must be executed if the condition part of the rule matches.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/action-protectionruleactiontype
import "encoding/xml"

type ActionProtectionRuleActionType struct {
	XMLName xml.Name

	// The Argument element specifies arguments to the action.
	Argument *Argument `xml:"Argument"`
	// Identifies the name of the action.
	Name *string `xml:"Name,attr"`
}

func (A *ActionProtectionRuleActionType) SetForMarshal() {
	A.XMLName.Local = "t:Action"
}

func (A *ActionProtectionRuleActionType) GetSchema() *Schema {
	return &SchemaTypes
}
