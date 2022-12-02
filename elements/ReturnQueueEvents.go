package elements

// The ReturnQueueEvents element indicates that the person who is running the task is in a privileged role.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/returnqueueevents
type ReturnQueueEvents bool

const (
	// false
	ReturnQueueEventsfalse ReturnQueueEvents = false
	// true
	ReturnQueueEventstrue ReturnQueueEvents = true
)
