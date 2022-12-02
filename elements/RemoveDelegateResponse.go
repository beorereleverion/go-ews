package elements

// The RemoveDelegateResponse element contains the status and result of a RemoveDelegate operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removedelegateresponse
type RemoveDelegateResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The ResponseMessages element contains the response messages for an Exchange Web Services delegate management request.
	ResponseMessages *ResponseMessagesArrayOfDelegateUserResponseMessageType `xml:"m:ResponseMessages"`
}
