package elements

// The InboxRules element represents an array of rules in the user's mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/inboxrules
import "encoding/xml"

type InboxRules struct {
	XMLName xml.Name

	// The Rule element contains a single rule and represents a rule in a user's mailbox.
	Rule *RuleRuleType `xml:"Rule"`
}

func (I *InboxRules) SetForMarshal() {
	I.XMLName.Local = "m:InboxRules"
}

func (I *InboxRules) GetSchema() *Schema {
	return &SchemaMessages
}
