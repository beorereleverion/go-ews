package elements

// The NoEndRecurrence element describes the start date of an item recurrence pattern that does not have a defined end date.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/noendrecurrence
type NoEndRecurrence struct {
	// The StartDate element represents the start date of a recurring task or calendar item.
	StartDate *StartDateRecurrence `xml:"t:StartDate"`
}
