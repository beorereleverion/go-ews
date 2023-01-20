package elements

// The RoutingType element represents the routing protocol for the recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/routingtype-emailaddress
import "encoding/xml"

type RoutingTypeEmailAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RoutingTypeEmailAddress) SetForMarshal() {
	R.XMLName.Local = "t:RoutingType"
}

func (R *RoutingTypeEmailAddress) GetSchema() *Schema {
	return &SchemaTypes
}
