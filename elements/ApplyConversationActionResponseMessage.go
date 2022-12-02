package elements

// The ApplyConversationActionResponseMessage element contains the status and results of an ApplyConversationAction operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/applyconversationactionresponsemessage
type ApplyConversationActionResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Describes the status of the response.The following values are valid for this attribute:SuccessWarningError
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	ApplyConversationActionResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed.The following are examples of sources of warnings:The Exchange store is offline during the batch.Active Directory Domain Services (AD DS) is offline.Mailboxes were moved.The message database (MDB) is offline.A password is expired.A quota has been exceeded.
	ApplyConversationActionResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled.The following are examples of sources of errors:  Invalid attributes or elementsAttributes or elements that are out of rangeAn unknown tag  An attribute or element that is not valid in the contextAn unauthorized access attempt by any clientA server-side failure in response to a valid client-side callInformation about the error can be found in the ResponseCode and MessageText elements.
	ApplyConversationActionResponseMessageError = `Error`
)
