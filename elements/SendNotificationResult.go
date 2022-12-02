package elements

// The SendNotificationResult element contains the response of a client application to a push notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendnotificationresult
type SendNotificationResult struct {
	// The SubscriptionStatus element describes the status of a push subscription.
	SubscriptionStatus *SubscriptionStatus `xml:"m:SubscriptionStatus"`
}
