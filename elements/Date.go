package elements

// The Date element represents the date that contains the suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/date
import "time"

type Date time.Time
