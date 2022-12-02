package elements

// The Timestamp element represents the timestamp of a mailbox event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timestamp
import "time"

type TimeStamp time.Time
