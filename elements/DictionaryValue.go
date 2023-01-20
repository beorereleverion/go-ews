package elements

// The DictionaryValue element specifies the dictionary value for a dictionary property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionaryvalue
import "encoding/xml"

type DictionaryValue struct {
	XMLName xml.Name

	// The Type element specifies a dictionary object type.
	Type *TypeUserConfiguration `xml:"Type"`
	// The Value element specifies the dictionary object value as a string.
	Value *ValueUserConfiguration `xml:"Value"`
}

func (D *DictionaryValue) SetForMarshal() {
	D.XMLName.Local = "t:DictionaryValue"
}

func (D *DictionaryValue) GetSchema() *Schema {
	return &SchemaTypes
}
