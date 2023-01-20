package elements

// The ReturnQueueEvents element indicates that the person who is running the task is in a privileged role.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/returnqueueevents
import "encoding/xml"

type ReturnQueueEvents struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

const (
	// false
	ReturnQueueEventsfalse bool = false
	// true
	ReturnQueueEventstrue bool = true
)

func (R *ReturnQueueEvents) SetForMarshal() {
	R.XMLName.Local = "m:ReturnQueueEvents"
}

func (R *ReturnQueueEvents) GetSchema() *Schema {
	return &SchemaMessages
}
