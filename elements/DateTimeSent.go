package elements

// The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimesent
import "time"

type DateTimeSent time.Time
