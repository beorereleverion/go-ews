package elements

// The SubscriptionId element represents the identifier for a streaming subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionid-getstreamingevents
import "encoding/xml"

type SubscriptionIdGetStreamingEvents struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SubscriptionIdGetStreamingEvents) SetForMarshal() {
	S.XMLName.Local = "t:SubscriptionId"
}

func (S *SubscriptionIdGetStreamingEvents) GetSchema() *Schema {
	return &SchemaTypes
}
