package elements

// The Values element specifies the values in an array of persona properties associated with an attributions array.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/values-arrayofstringvaluetype
type ValuesArrayOfStringValueType struct {
	// The Value element specifies a value in an array of persona properties associated with an attributions array.
	Value *ValueArrayOfStringValueType `xml:"t:Value"`
}
