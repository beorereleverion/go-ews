package elements

// The WorkingPeriodArray element contains working period information for the mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workingperiodarray
type WorkingPeriodArray struct {
	// The WorkingPeriod element contains the work week days and hours of the mailbox user.
	WorkingPeriod *WorkingPeriod `xml:"t:WorkingPeriod"`
}
