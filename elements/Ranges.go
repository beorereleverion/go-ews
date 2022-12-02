package elements

// The Ranges element specifies an array of recurrence ranges.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ranges
type Ranges struct {
	// The Range element specifies a range of calendar item occurrences for a repeating calendar item.
	Range *Range `xml:"t:Range"`
}
