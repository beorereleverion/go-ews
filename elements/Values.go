package elements

// The Values element contains a collection of values for an extended property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/values
type Values struct {
	// The Value element contains the value of an extended property.
	Value *Value `xml:"t:Value"`
}
