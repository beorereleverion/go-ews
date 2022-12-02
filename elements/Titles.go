package elements

// The Titles element specifies an array of job titles and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/titles
type Titles struct {
	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"t:StringAttributedValue"`
}
