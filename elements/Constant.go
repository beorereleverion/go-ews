package elements

// The Constant element identifies a constant value in a restriction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/constant
type Constant struct {
	// Specifies the value to compare in the restriction.
	Value *string `xml:"Value,attr"`
}
