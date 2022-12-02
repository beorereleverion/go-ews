package elements

// The SubmitTime element represents the time at which the message was sent to the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/submittime
import "time"

type SubmitTime time.Time
