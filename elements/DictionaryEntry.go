package elements

// The DictionaryEntry element specifies the contents of a single dictionary entry property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionaryentry
type DictionaryEntry struct {
	// The DictionaryKey element specifies the dictionary key for a dictionary property.
	DictionaryKey *DictionaryKey `xml:"t:DictionaryKey"`
	// The DictionaryValue element specifies the dictionary value for a dictionary property.
	DictionaryValue *DictionaryValue `xml:"t:DictionaryValue"`
}
