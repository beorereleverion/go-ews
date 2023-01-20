package elements

// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stringattributedvalue
import "encoding/xml"

type StringAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (S *StringAttributedValue) SetForMarshal() {
	S.XMLName.Local = "t:StringAttributedValue"
}

func (S *StringAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
