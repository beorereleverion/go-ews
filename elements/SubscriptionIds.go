package elements

// The SubscriptionIds element contains an array of subscription identifiers that identify the subscriptions to get streaming events for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionids
import "encoding/xml"

type SubscriptionIds struct {
	XMLName xml.Name

	// The SubscriptionId element represents the identifier for a streaming subscription.
	SubscriptionId *SubscriptionIdGetStreamingEvents `xml:"SubscriptionId"`
}

func (S *SubscriptionIds) SetForMarshal() {
	S.XMLName.Local = "m:SubscriptionIds"
}

func (S *SubscriptionIds) GetSchema() *Schema {
	return &SchemaMessages
}
