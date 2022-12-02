package elements

// The ModifiedOccurrences element contains an array of recurring calendar item occurrences that have been modified so that they are different than the recurrence master item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/modifiedoccurrences
type ModifiedOccurrences struct {
	// The Occurrence element represents a single modified occurrence of a recurring calendar item.
	Occurrence *Occurrence `xml:"t:Occurrence"`
}
