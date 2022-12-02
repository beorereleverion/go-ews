package elements

// The ReminderGroup element specifies whether the reminder is for a calendar item or a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindergroup
type ReminderGroup string

const (
	// Calendar
	ReminderGroupCalendar ReminderGroup = `Calendar`
	// Task
	ReminderGroupTask ReminderGroup = `Task`
)
