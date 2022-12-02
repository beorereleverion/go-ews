package elements

// The DeletedOccurrence element represents a deleted occurrence of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletedoccurrence
type DeletedOccurrence struct {
	// The Start element represents the start of a duration.
	Start *Start `xml:"t:Start"`
}
