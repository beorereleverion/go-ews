package elements

// The EventTypes element contains a collection of event notification types that are used to create a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventtypes
type EventTypes struct {
	// The EventType element is used to create a subscription and identifies an event type to be reported in a notification.
	EventType *EventType `xml:"t:EventType"`
}
