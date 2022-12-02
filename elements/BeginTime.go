package elements

// The BeginTime element specifies the beginning of the time span to query for reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/begintime
import "time"

type BeginTime time.Time
