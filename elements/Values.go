package elements

// The Values element contains a collection of values for an extended property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/values
import "encoding/xml"

type Values struct {
	XMLName xml.Name

	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (V *Values) SetForMarshal() {
	V.XMLName.Local = "t:Values"
}

func (V *Values) GetSchema() *Schema {
	return &SchemaTypes
}
