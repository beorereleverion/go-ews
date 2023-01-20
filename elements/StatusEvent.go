package elements

// The StatusEvent element represents a notification that no new activity has occurred in the mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/statusevent
import "encoding/xml"

type StatusEvent struct {
	XMLName xml.Name

	// The Watermark element represents an event bookmark in the mailbox event queue.
	Watermark *Watermark `xml:"Watermark"`
}

func (S *StatusEvent) SetForMarshal() {
	S.XMLName.Local = "t:StatusEvent"
}

func (S *StatusEvent) GetSchema() *Schema {
	return &SchemaTypes
}
