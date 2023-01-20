package elements

// The SubscriptionStatus element describes the status of a push subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionstatus
import "encoding/xml"

type SubscriptionStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SubscriptionStatus) SetForMarshal() {
	S.XMLName.Local = "m:SubscriptionStatus"
}

func (S *SubscriptionStatus) GetSchema() *Schema {
	return &SchemaMessages
}
