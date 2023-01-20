package elements

// The RuleOperationError element represents a rule operation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ruleoperationerror
import "encoding/xml"

type RuleOperationError struct {
	XMLName xml.Name

	// The OperationIndex element specifies the index of the operation in the request that caused the rule operation error.
	OperationIndex *OperationIndex `xml:"OperationIndex"`
	// The ValidationErrors element represents an array of rule validation errors on each rule field that has an error.
	ValidationErrors *ValidationErrors `xml:"ValidationErrors"`
}

func (R *RuleOperationError) SetForMarshal() {
	R.XMLName.Local = "m:RuleOperationError"
}

func (R *RuleOperationError) GetSchema() *Schema {
	return &SchemaMessages
}
