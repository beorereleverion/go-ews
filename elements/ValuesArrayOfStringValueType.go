package elements

// The Values element specifies the values in an array of persona properties associated with an attributions array.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/values-arrayofstringvaluetype
import "encoding/xml"

type ValuesArrayOfStringValueType struct {
	XMLName xml.Name

	// The Value element specifies a value in an array of persona properties associated with an attributions array.
	Value *ValueArrayOfStringValueType `xml:"Value"`
}

func (V *ValuesArrayOfStringValueType) SetForMarshal() {
	V.XMLName.Local = "t:Values"
}

func (V *ValuesArrayOfStringValueType) GetSchema() *Schema {
	return &SchemaTypes
}
