package elements

// The RoutingType element defines the address type for the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/routingtype-emailaddresstype
import "encoding/xml"

type RoutingTypeEmailAddressType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RoutingTypeEmailAddressType) SetForMarshal() {
	R.XMLName.Local = "t:RoutingType"
}

func (R *RoutingTypeEmailAddressType) GetSchema() *Schema {
	return &SchemaTypes
}
