package elements

// The ReminderNextTime element specifies the date and time for the next reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindernexttime
import "time"

type ReminderNextTime time.Time
