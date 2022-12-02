package elements

// The CompleteDate element represents the date on which an item was completed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/completedate
import "time"

type CompleteDate time.Time
