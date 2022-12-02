package elements

// The AssignedTime element represents the time when a task is assigned to a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assignedtime
import "time"

type AssignedTime time.Time
