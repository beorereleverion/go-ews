package elements

// The DictionaryKey element specifies the dictionary key for a dictionary property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionarykey
import "encoding/xml"

type DictionaryKey struct {
	XMLName xml.Name

	// The Type element specifies a dictionary object type.
	Type *TypeUserConfiguration `xml:"Type"`
	// The Value element specifies the dictionary object value as a string.
	Value *ValueUserConfiguration `xml:"Value"`
}

func (D *DictionaryKey) SetForMarshal() {
	D.XMLName.Local = "t:DictionaryKey"
}

func (D *DictionaryKey) GetSchema() *Schema {
	return &SchemaTypes
}
