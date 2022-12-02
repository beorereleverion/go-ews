package elements

// The Date element represents the date and time at which the event occurred.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/date-messagetracking
import "time"

type DateMessageTracking time.Time
