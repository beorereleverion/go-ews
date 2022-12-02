package elements

// The Managers element specifies an array of manager names and the identifiers of their source attributions for a persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/managers
type Managers struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
