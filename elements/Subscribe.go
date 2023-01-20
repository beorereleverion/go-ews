package elements

// The Subscribe element contains the properties used to create subscriptions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscribe
import "encoding/xml"

type Subscribe struct {
	XMLName xml.Name

	// The PullSubscriptionRequest element represents a subscription to a pull-based event notification subscription.
	PullSubscriptionRequest *PullSubscriptionRequest `xml:"PullSubscriptionRequest"`
	// The PushSubscriptionRequest element represents a subscription to a push-based event notification subscription.
	PushSubscriptionRequest *PushSubscriptionRequest `xml:"PushSubscriptionRequest"`
	// The StreamingSubscriptionRequest element represents a subscription to a streaming event notification subscription.
	StreamingSubscriptionRequest *StreamingSubscriptionRequest `xml:"StreamingSubscriptionRequest"`
}

func (S *Subscribe) SetForMarshal() {
	S.XMLName.Local = "m:Subscribe"
}

func (S *Subscribe) GetSchema() *Schema {
	return &SchemaMessages
}
