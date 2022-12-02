package elements

// The NumberedRecurrence element describes the start date and the number of occurrences of a recurring item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberedrecurrence
type NumberedRecurrence struct {
	// The NumberOfOccurrences element contains the number of occurrences of a recurring item.
	NumberOfOccurrences *NumberOfOccurrences `xml:"t:NumberOfOccurrences"`
	// The StartDate element represents the start date of a recurring task or calendar item.
	StartDate *StartDateRecurrence `xml:"t:StartDate"`
}
