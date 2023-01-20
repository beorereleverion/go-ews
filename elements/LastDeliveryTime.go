package elements

// The LastDeliveryTime element contains the delivery time of the message that was last received in this conversation in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastdeliverytime
import "time"

import "encoding/xml"

type LastDeliveryTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (L *LastDeliveryTime) SetForMarshal() {
	L.XMLName.Local = "t:LastDeliveryTime"
}

func (L *LastDeliveryTime) GetSchema() *Schema {
	return &SchemaTypes
}
