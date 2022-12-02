package elements

// The GetEvents element represents the operation used by pull clients to request notifications from the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getevents
type GetEvents struct {
	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"t:Watermark"`
}
