package elements

// The StopProcessingRules element indicates whether subsequent rules are to be evaluated.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stopprocessingrules
import "encoding/xml"

type StopProcessingRules struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	StopProcessingRulesfalse bool = false
	// true
	StopProcessingRulestrue bool = true
)

func (S *StopProcessingRules) SetForMarshal() {
	S.XMLName.Local = "m:StopProcessingRules"
}

func (S *StopProcessingRules) GetSchema() *Schema {
	return &SchemaMessages
}
