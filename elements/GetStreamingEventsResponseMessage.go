package elements

// The GetStreamingEventsResponseMessage element contains the status and result of a single GetStreamingEvents operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getstreamingeventsresponsemessage
type GetStreamingEventsResponseMessage struct {
	// The ConnectionStatus element provides a text description of the status of a streaming subscription.
	ConnectionStatus *ConnectionStatus `xml:"m:ConnectionStatus"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The ErrorSubscriptionIds element contains an array of invalid subscription IDs.
	ErrorSubscriptionIds *ErrorSubscriptionIds `xml:"m:ErrorSubscriptionIds"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The Notifications element contains an array of information about the subscription and the events that have occurred since the last notification.
	Notifications *Notifications `xml:"m:Notifications"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Describes the status of a GetStreamingEvents operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}
