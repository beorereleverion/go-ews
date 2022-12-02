package elements

// The LastAttemptTime element contains the time and date at which the last attempt to index the item was made.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastattempttime
import "time"

type LastAttemptTime time.Time
