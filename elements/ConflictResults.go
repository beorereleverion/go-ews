package elements

// The ConflictResults element contains the number of conflicts in an UpdateItem operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conflictresults
type ConflictResults struct {
	// The Count element contains the number of conflicts in an UpdateItem operation response.
	Count *Count `xml:"t:Count"`
}
