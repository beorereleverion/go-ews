package elements

// The EndDate element specifies the end date of the item the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddate-remindertype
import "time"

type EndDateReminderType time.Time
