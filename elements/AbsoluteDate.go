package elements

// The AbsoluteDate element represents the date when the time changes from standard or daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/absolutedate
import "time"

type AbsoluteDate time.Time
