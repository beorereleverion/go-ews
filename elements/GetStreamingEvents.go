package elements

// The GetStreamingEvents element represents the operation that is used by clients to request streaming notifications from the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getstreamingevents
import "encoding/xml"

type GetStreamingEvents struct {
	XMLName xml.Name

	// The ConnectionTimeout element specifies the number of minutes to keep a connection open.
	ConnectionTimeout *ConnectionTimeout `xml:"ConnectionTimeout"`
	// The SubscriptionIds element contains an array of subscription identifiers that identify the subscriptions to get streaming events for.
	SubscriptionIds *SubscriptionIds `xml:"SubscriptionIds"`
}

func (G *GetStreamingEvents) SetForMarshal() {
	G.XMLName.Local = "m:GetStreamingEvents"
}

func (G *GetStreamingEvents) GetSchema() *Schema {
	return &SchemaMessages
}
