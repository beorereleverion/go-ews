package elements

// The ValidationErrors element represents an array of rule validation errors on each rule field that has an error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/validationerrors
type ValidationErrors struct {
	// The Error element represents a single validation error on a particular rule property value, predicate property value, or action property value.
	Error *Error `xml:"m:Error"`
}
