package elements

// The Operations element contains an array of rule operations that can be performed on an Inbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/operations
import "encoding/xml"

type Operations struct {
	XMLName xml.Name

	// The CreateRuleOperation element represents an operation to create a new Inbox rule.
	CreateRuleOperation *CreateRuleOperation `xml:"CreateRuleOperation"`
	// The DeleteRuleOperation element contains an operation to delete an existing Inbox rule.
	DeleteRuleOperation *DeleteRuleOperation `xml:"DeleteRuleOperation"`
	// The SetRuleOperation element represents an operation to update an existing rule.
	SetRuleOperation *SetRuleOperation `xml:"SetRuleOperation"`
}

func (O *Operations) SetForMarshal() {
	O.XMLName.Local = "t:Operations"
}

func (O *Operations) GetSchema() *Schema {
	return &SchemaTypes
}
