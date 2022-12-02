package elements

// The FromDate element specifies the date that the message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fromdate
import "time"

type FromDate time.Time
