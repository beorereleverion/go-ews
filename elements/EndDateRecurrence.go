package elements

// The EndDate element represents the end date of a recurring task or a calendar item that has the EndDateRecurrence pattern type.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddate-recurrence
import "time"

type EndDateRecurrence time.Time
