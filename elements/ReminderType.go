package elements

// The ReminderType element specifies the type of reminders to return.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindertype
type ReminderType string

const (
	// All
	ReminderTypeAll ReminderType = `All`
	// Current
	ReminderTypeCurrent ReminderType = `Current`
	// Old
	ReminderTypeOld ReminderType = `Old`
)
