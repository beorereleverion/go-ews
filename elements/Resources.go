package elements

// The Resources element represents a scheduled resource for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resources
import "encoding/xml"

type Resources struct {
	XMLName xml.Name

	// The Attendee element represents attendees and resources for a meeting.
	Attendee *Attendee `xml:"Attendee"`
}

func (R *Resources) SetForMarshal() {
	R.XMLName.Local = "t:Resources"
}

func (R *Resources) GetSchema() *Schema {
	return &SchemaTypes
}
