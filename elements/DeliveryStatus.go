package elements

// The DeliveryStatus element specifies the status for a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deliverystatus
import "encoding/xml"

type DeliveryStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Delivered
	DeliveryStatusDelivered string = `Delivered`
	// Pending
	DeliveryStatusPending string = `Pending`
	// Read
	DeliveryStatusRead string = `Read`
	// Transferred
	DeliveryStatusTransferred string = `Transferred`
	// Unsuccessful
	DeliveryStatusUnsuccessful string = `Unsuccessful`
)

func (D *DeliveryStatus) SetForMarshal() {
	D.XMLName.Local = "t:DeliveryStatus"
}

func (D *DeliveryStatus) GetSchema() *Schema {
	return &SchemaTypes
}
