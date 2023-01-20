package elements

// The GetEvents element represents the operation used by pull clients to request notifications from the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getevents
import "encoding/xml"

type GetEvents struct {
	XMLName xml.Name

	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"Watermark"`
}

func (G *GetEvents) SetForMarshal() {
	G.XMLName.Local = "m:GetEvents"
}

func (G *GetEvents) GetSchema() *Schema {
	return &SchemaMessages
}
