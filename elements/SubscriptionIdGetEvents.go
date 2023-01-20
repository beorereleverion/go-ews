package elements

// The SubscriptionId element represents the identifier for a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/subscriptionid-getevents
import "encoding/xml"

type SubscriptionIdGetEvents struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
