package elements

// The ErrorSubscriptionIds element contains an array of invalid subscription IDs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/errorsubscriptionids
import "encoding/xml"

type ErrorSubscriptionIds struct {
	XMLName xml.Name

	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
}

func (E *ErrorSubscriptionIds) SetForMarshal() {
	E.XMLName.Local = "m:ErrorSubscriptionIds"
}

func (E *ErrorSubscriptionIds) GetSchema() *Schema {
	return &SchemaMessages
}
