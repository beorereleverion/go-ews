package elements

// The ReminderTime element specifies the time for the reminder to occur.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindertime
import "time"

type ReminderTime time.Time
