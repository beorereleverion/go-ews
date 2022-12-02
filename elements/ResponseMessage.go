package elements

// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessage
type ResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Represents the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	ResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are some possible causes for warnings:  - The Exchange store is offline during the batch.  - The Active Directory directory service is offline.  - Mailboxes are moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota is exceeded.
	ResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are some possible causes for errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempt by any client  - Server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	ResponseMessageError = `Error`
)
