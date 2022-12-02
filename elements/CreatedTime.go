package elements

// The CreatedTime element specifies the time at which the item was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createdtime
import "time"

type CreatedTime time.Time
