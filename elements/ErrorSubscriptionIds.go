package elements

// The ErrorSubscriptionIds element contains an array of invalid subscription IDs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errorsubscriptionids
type ErrorSubscriptionIds struct {
	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
}
