package elements

// The Argument element specifies arguments to the action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/argument
type Argument struct {
	// A non-empty string value that represents the value of an argument to the action part of a protection rule. This attribute is required.
	Value *string `xml:"Value,attr"`
}
