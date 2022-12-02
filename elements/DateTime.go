package elements

// The DateTime element represents the date and time at which the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetime
import "time"

type DateTime time.Time
