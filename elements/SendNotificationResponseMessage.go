package elements

// The SendNotificationResponseMessage element contains the status and result of a single SendNotification operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendnotificationresponsemessage
import "encoding/xml"

type SendNotificationResponseMessage struct {
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
	// Describes the status of a SendNotification operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (S *SendNotificationResponseMessage) SetForMarshal() {
	S.XMLName.Local = "m:SendNotificationResponseMessage"
}

func (S *SendNotificationResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
