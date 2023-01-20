package elements

// The Dictionary element defines a set of dictionary property entries for a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionary
import "encoding/xml"

type Dictionary struct {
	XMLName xml.Name

	// The DictionaryEntry element specifies the contents of a single dictionary entry property.
	DictionaryEntry *DictionaryEntry `xml:"DictionaryEntry"`
}

func (D *Dictionary) SetForMarshal() {
	D.XMLName.Local = "t:Dictionary"
}

func (D *Dictionary) GetSchema() *Schema {
	return &SchemaTypes
}
