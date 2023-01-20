package elements

// The SenderDepartments element specifies that the department of the sender matches any of the specified departments in the child Value (ProtectionRuleValueType) elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senderdepartments
import "encoding/xml"

type SenderDepartments struct {
	XMLName xml.Name

	// The Value element identifies a single recipient or sender department.
	Value *ValueProtectionRuleValueType `xml:"Value"`
}

func (S *SenderDepartments) SetForMarshal() {
	S.XMLName.Local = "t:SenderDepartments"
}

func (S *SenderDepartments) GetSchema() *Schema {
	return &SchemaTypes
}
