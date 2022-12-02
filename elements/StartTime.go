package elements

// The StartTime element represents the start of a time span.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime
import "time"

type StartTime time.Time
