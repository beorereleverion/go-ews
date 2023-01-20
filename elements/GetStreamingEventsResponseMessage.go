package elements

// The GetStreamingEventsResponseMessage element contains the status and result of a single GetStreamingEvents operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getstreamingeventsresponsemessage
import "encoding/xml"

type GetStreamingEventsResponseMessage struct {
	XMLName xml.Name

	// The ConnectionStatus element provides a text description of the status of a streaming subscription.
	ConnectionStatus *ConnectionStatus `xml:"ConnectionStatus"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The ErrorSubscriptionIds element contains an array of invalid subscription IDs.
	ErrorSubscriptionIds *ErrorSubscriptionIds `xml:"ErrorSubscriptionIds"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The Notifications element contains an array of information about the subscription and the events that have occurred since the last notification.
	Notifications *Notifications `xml:"Notifications"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a GetStreamingEvents operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetStreamingEventsResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetStreamingEventsResponseMessage"
}

func (G *GetStreamingEventsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
