package elements

// The StartDate element represents the start date of a recurring task or calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdate-recurrence
import "time"

type StartDateRecurrence time.Time
