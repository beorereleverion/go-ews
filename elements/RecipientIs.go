package elements

// The RecipientIs element specifies that any recipient of the e-mail message matches any of the specified recipients in the child Value (ProtectionRuleValueType) elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipientis
import "encoding/xml"

type RecipientIs struct {
	XMLName xml.Name

	// The Value element identifies a single recipient or sender department.
	Value *ValueProtectionRuleValueType `xml:"Value"`
}

func (R *RecipientIs) SetForMarshal() {
	R.XMLName.Local = "t:RecipientIs"
}

func (R *RecipientIs) GetSchema() *Schema {
	return &SchemaTypes
}
