package elements

// The DeletedOccurrences element contains an array of deleted occurrences of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedoccurrences
type DeletedOccurrences struct {
	// The DeletedOccurrence element represents a deleted occurrence of a recurring calendar item.
	DeletedOccurrence *DeletedOccurrence `xml:"t:DeletedOccurrence"`
}
