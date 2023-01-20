package elements

// The RecipientTrackingEvent element contains information for a single event for a recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipienttrackingevent
import "encoding/xml"

type RecipientTrackingEvent struct {
	XMLName xml.Name

	// The BccRecipient element represents a recipient to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipient *BccRecipient `xml:"BccRecipient"`
	// The Date element represents the date and time at which the event occurred.
	Date *DateMessageTracking `xml:"Date"`
	// The DeliveryStatus element specifies the status for a message.
	DeliveryStatus *DeliveryStatus `xml:"DeliveryStatus"`
	// The EventData element represents data that is associated with the processing step for the event.
	EventData *EventData `xml:"EventData"`
	// The EventDescription element
	EventDescription *EventDescription `xml:"EventDescription"`
	// The HiddenRecipient element indicates that the recipient was added by an organization policy that should be hidden from unprivileged users.
	HiddenRecipient *HiddenRecipient `xml:"HiddenRecipient"`
	// The InternalId element represents an integer value for the event identification.
	InternalId *InternalId `xml:"InternalId"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The Recipient element represents the recipient for whom the event occurred.
	Recipient *Recipient `xml:"Recipient"`
	// The RootAddress element represents the first address that starts the event for a RecipientTrackingEvent event.
	RootAddress *RootAddress `xml:"RootAddress"`
	// The Server element represents the physical server where the event occurred.
	Server *ServerMessageTracking `xml:"Server"`
	// The UniquePathId element represents a string that is different for each path in a tracking report.
	UniquePathId *UniquePathId `xml:"UniquePathId"`
}

func (R *RecipientTrackingEvent) SetForMarshal() {
	R.XMLName.Local = "t:RecipientTrackingEvent"
}

func (R *RecipientTrackingEvent) GetSchema() *Schema {
	return &SchemaTypes
}
