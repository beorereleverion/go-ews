package elements

// The EventData element represents data that is associated with the processing step for the event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/eventdata
import "encoding/xml"

type EventData struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}

func (E *EventData) SetForMarshal() {
	E.XMLName.Local = "t:EventData"
}

func (E *EventData) GetSchema() *Schema {
	return &SchemaTypes
}
