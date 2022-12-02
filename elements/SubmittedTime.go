package elements

// The SubmittedTime element represents the time that the message entered the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/submittedtime
import "time"

type SubmittedTime time.Time
