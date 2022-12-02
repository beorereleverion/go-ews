package elements

// The Ids element contains an array of time zone definition identifiers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ids
type Ids struct {
	// The Id element identifies a single time zone definition.
	Id *IdTimeZone `xml:"t:Id"`
}
