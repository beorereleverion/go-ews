package elements

// The RecipientTrackingEvent element contains information for a single event for a recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recipienttrackingevent
type RecipientTrackingEvent struct {
	// The BccRecipient element represents a recipient to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipient *BccRecipient `xml:"t:BccRecipient"`
	// The Date element represents the date and time at which the event occurred.
	Date *DateMessageTracking `xml:"t:Date"`
	// The DeliveryStatus element specifies the status for a message.
	DeliveryStatus *DeliveryStatus `xml:"t:DeliveryStatus"`
	// The EventData element represents data that is associated with the processing step for the event.
	EventData *EventData `xml:"t:EventData"`
	// The EventDescription element
	EventDescription *EventDescription `xml:"t:EventDescription"`
	// The HiddenRecipient element indicates that the recipient was added by an organization policy that should be hidden from unprivileged users.
	HiddenRecipient *HiddenRecipient `xml:"t:HiddenRecipient"`
	// The InternalId element represents an integer value for the event identification.
	InternalId *InternalId `xml:"t:InternalId"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"m:Properties"`
	// The Recipient element represents the recipient for whom the event occurred.
	Recipient *Recipient `xml:"t:Recipient"`
	// The RootAddress element represents the first address that starts the event for a RecipientTrackingEvent event.
	RootAddress *RootAddress `xml:"t:RootAddress"`
	// The Server element represents the physical server where the event occurred.
	Server *ServerMessageTracking `xml:"t:Server"`
	// The UniquePathId element represents a string that is different for each path in a tracking report.
	UniquePathId *UniquePathId `xml:"t:UniquePathId"`
}
