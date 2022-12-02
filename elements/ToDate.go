package elements

// The ToDate element specifies the date that the message was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/todate
import "time"

type ToDate time.Time
