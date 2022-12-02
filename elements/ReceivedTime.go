package elements

// The ReceivedTime element specifies the time at which an item was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedtime
import "time"

type ReceivedTime time.Time
