package elements

// The GetEventsResponseMessage element contains the status and result of a single GetEvents operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/geteventsresponsemessage
import "encoding/xml"

type GetEventsResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The Notification element contains information about the subscription and the events that have occurred since the last notification.
	Notification *Notification `xml:"Notification"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of a GetEvents operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetEventsResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetEventsResponseMessage"
}

func (G *GetEventsResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
