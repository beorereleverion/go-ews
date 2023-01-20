package elements

// The StringArrayAttributedValue element specifies an instance of an array of string data for a persona element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/stringarrayattributedvalue
import "encoding/xml"

type StringArrayAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (S *StringArrayAttributedValue) SetForMarshal() {
	S.XMLName.Local = "t:StringArrayAttributedValue"
}

func (S *StringArrayAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
