package elements

// The RecipientTrackingEvents element represents a collection of one or more events for a message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipienttrackingevents
import "encoding/xml"

type RecipientTrackingEvents struct {
	XMLName xml.Name

	// The RecipientTrackingEvent element contains information for a single event for a recipient.
	RecipientTrackingEvent *RecipientTrackingEvent `xml:"RecipientTrackingEvent"`
}

func (R *RecipientTrackingEvents) SetForMarshal() {
	R.XMLName.Local = "t:RecipientTrackingEvents"
}

func (R *RecipientTrackingEvents) GetSchema() *Schema {
	return &SchemaTypes
}
