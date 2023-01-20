package elements

// The EventTypes element contains a collection of event notification types that are used to create a subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventtypes
import "encoding/xml"

type EventTypes struct {
	XMLName xml.Name

	// The EventType element is used to create a subscription and identifies an event type to be reported in a notification.
	EventType *EventType `xml:"EventType"`
}

func (E *EventTypes) SetForMarshal() {
	E.XMLName.Local = "t:EventTypes"
}

func (E *EventTypes) GetSchema() *Schema {
	return &SchemaTypes
}
