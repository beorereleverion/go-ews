package elements

// The StatusEvent element represents a notification that no new activity has occurred in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/statusevent
type StatusEvent struct {
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
}
