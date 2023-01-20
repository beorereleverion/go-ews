package elements

// The GlobalLastDeliveryTime element contains the delivery time of the message that was last received in this conversation across all folders in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globallastdeliverytime
import "time"

import "encoding/xml"

type GlobalLastDeliveryTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (G *GlobalLastDeliveryTime) SetForMarshal() {
	G.XMLName.Local = "t:GlobalLastDeliveryTime"
}

func (G *GlobalLastDeliveryTime) GetSchema() *Schema {
	return &SchemaTypes
}
