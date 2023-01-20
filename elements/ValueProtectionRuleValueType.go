package elements

// The Value element identifies a single recipient or sender department.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-protectionrulevaluetype
import "encoding/xml"

type ValueProtectionRuleValueType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *ValueProtectionRuleValueType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueProtectionRuleValueType) GetSchema() *Schema {
	return &SchemaTypes
}
