package elements

// The DueDate element represents the date an item is due.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/duedate
import "time"

type DueDate time.Time
