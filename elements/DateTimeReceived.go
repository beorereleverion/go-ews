package elements

// The DateTimeReceived element represents the date and time that an item in a mailbox was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimereceived
import "time"

type DateTimeReceived time.Time
