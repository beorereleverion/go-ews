package elements

// The Subscribe element contains the properties used to create subscriptions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscribe
type Subscribe struct {
	// The PullSubscriptionRequest element represents a subscription to a pull-based event notification subscription.
	PullSubscriptionRequest *PullSubscriptionRequest `xml:"m:PullSubscriptionRequest"`
	// The PushSubscriptionRequest element represents a subscription to a push-based event notification subscription.
	PushSubscriptionRequest *PushSubscriptionRequest `xml:"m:PushSubscriptionRequest"`
	// The StreamingSubscriptionRequest element represents a subscription to a streaming event notification subscription.
	StreamingSubscriptionRequest *StreamingSubscriptionRequest `xml:"m:StreamingSubscriptionRequest"`
}
