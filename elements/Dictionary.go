package elements

// The Dictionary element defines a set of dictionary property entries for a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionary
type Dictionary struct {
	// The DictionaryEntry element specifies the contents of a single dictionary entry property.
	DictionaryEntry *DictionaryEntry `xml:"t:DictionaryEntry"`
}
