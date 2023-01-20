package elements

// The DeliveryRestricted element indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deliveryrestricted
import "encoding/xml"

type DeliveryRestricted struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	DeliveryRestrictedfalse bool = false
	// true
	DeliveryRestrictedtrue bool = true
)

func (D *DeliveryRestricted) SetForMarshal() {
	D.XMLName.Local = "t:DeliveryRestricted"
}

func (D *DeliveryRestricted) GetSchema() *Schema {
	return &SchemaTypes
}
