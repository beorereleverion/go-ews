package elements

// The IsReminderSet element indicates whether a reminder has been set for the calendar event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isreminderset
type IsReminderSet bool

const (
	// false
	IsReminderSetfalse IsReminderSet = false
	// true
	IsReminderSettrue IsReminderSet = true
)
