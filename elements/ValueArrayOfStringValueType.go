package elements

// The Value element specifies a value in an array of persona properties associated with an attributions array.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-arrayofstringvaluetype
import "encoding/xml"

type ValueArrayOfStringValueType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (V *ValueArrayOfStringValueType) SetForMarshal() {
	V.XMLName.Local = "t:Value"
}

func (V *ValueArrayOfStringValueType) GetSchema() *Schema {
	return &SchemaTypes
}
