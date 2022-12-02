package elements

// The SentTime element specifies the time at which the item was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttime
import "time"

type SentTime time.Time
