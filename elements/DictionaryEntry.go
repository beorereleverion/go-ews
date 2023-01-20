package elements

// The DictionaryEntry element specifies the contents of a single dictionary entry property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionaryentry
import "encoding/xml"

type DictionaryEntry struct {
	XMLName xml.Name

	// The DictionaryKey element specifies the dictionary key for a dictionary property.
	DictionaryKey *DictionaryKey `xml:"DictionaryKey"`
	// The DictionaryValue element specifies the dictionary value for a dictionary property.
	DictionaryValue *DictionaryValue `xml:"DictionaryValue"`
}

func (D *DictionaryEntry) SetForMarshal() {
	D.XMLName.Local = "t:DictionaryEntry"
}

func (D *DictionaryEntry) GetSchema() *Schema {
	return &SchemaTypes
}
