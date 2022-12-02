package elements

// The DictionaryKey element specifies the dictionary key for a dictionary property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/dictionarykey
type DictionaryKey struct {
	// The Type element specifies a dictionary object type.
	Type *TypeUserConfiguration `xml:"t:Type"`
	// The Value element specifies the dictionary object value as a string.
	Value *ValueUserConfiguration `xml:"t:Value"`
}
