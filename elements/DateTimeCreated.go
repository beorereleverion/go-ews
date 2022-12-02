package elements

// The DateTimeCreated element represents the date and time that an item in the mailbox was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimecreated
import "time"

type DateTimeCreated time.Time
