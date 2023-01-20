package elements

// The RuleId element specifies a rule identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ruleid
import "encoding/xml"

type RuleId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RuleId) SetForMarshal() {
	R.XMLName.Local = "m:RuleId"
}

func (R *RuleId) GetSchema() *Schema {
	return &SchemaMessages
}
