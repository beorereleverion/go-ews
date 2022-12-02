package elements

// The RecurrenceId element is used to identify a specific instance of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurrenceid
import "time"

type RecurrenceId time.Time
