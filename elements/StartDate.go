package elements

// The StartDate element represents the start date of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdate
import "time"

type StartDate time.Time
