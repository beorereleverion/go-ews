package elements

// The EventType element is used to create a subscription and identifies an event type to be reported in a notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventtype
import "encoding/xml"

type EventType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EventType) SetForMarshal() {
	E.XMLName.Local = "t:EventType"
}

func (E *EventType) GetSchema() *Schema {
	return &SchemaTypes
}
