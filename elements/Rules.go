package elements

// The Rules element contains an array of protection rules.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rules-ex15websvcsotherref
import "encoding/xml"

type Rules struct {
	XMLName xml.Name

	// The Rule element contains a single protection rule.
	Rule *Rule `xml:"Rule"`
}

func (R *Rules) SetForMarshal() {
	R.XMLName.Local = "t:Rules"
}

func (R *Rules) GetSchema() *Schema {
	return &SchemaTypes
}
