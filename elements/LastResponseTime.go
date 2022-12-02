package elements

// The LastResponseTime element represents the date and time of the latest response received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastresponsetime
import "time"

type LastResponseTime time.Time
