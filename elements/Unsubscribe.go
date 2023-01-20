package elements

// The Unsubscribe element contains the properties used to unsubscribe from a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unsubscribe
import "encoding/xml"

type Unsubscribe struct {
	XMLName xml.Name

	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
}

func (U *Unsubscribe) SetForMarshal() {
	U.XMLName.Local = "m:Unsubscribe"
}

func (U *Unsubscribe) GetSchema() *Schema {
	return &SchemaMessages
}
