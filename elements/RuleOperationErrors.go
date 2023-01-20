package elements

// The RuleOperationErrors element represents an array of rule validation errors on each rule field that has an error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ruleoperationerrors
import "encoding/xml"

type RuleOperationErrors struct {
	XMLName xml.Name

	// The RuleOperationError element represents a rule operation error.
	RuleOperationError *RuleOperationError `xml:"RuleOperationError"`
}

func (R *RuleOperationErrors) SetForMarshal() {
	R.XMLName.Local = "m:RuleOperationErrors"
}

func (R *RuleOperationErrors) GetSchema() *Schema {
	return &SchemaMessages
}
