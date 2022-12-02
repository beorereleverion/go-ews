package elements

// The SubscriptionIds element contains an array of subscription identifiers that identify the subscriptions to get streaming events for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionids
type SubscriptionIds struct {
	// The SubscriptionId element represents the identifier for a streaming subscription.
	SubscriptionId *SubscriptionIdGetStreamingEvents `xml:"t:SubscriptionId"`
}
