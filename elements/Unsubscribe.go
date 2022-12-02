package elements

// The Unsubscribe element contains the properties used to unsubscribe from a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unsubscribe
type Unsubscribe struct {
	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
}
