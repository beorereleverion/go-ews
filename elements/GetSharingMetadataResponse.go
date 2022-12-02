package elements

// The GetSharingMetadataResponse element defines a response to a GetSharingMetadata operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsharingmetadataresponse
type GetSharingMetadataResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The EncryptedSharedFolderDataCollection element contains a collection of data structures that a client can use to authorize the sharing of its calendar or contact data with other clients.
	EncryptedSharedFolderDataCollection *EncryptedSharedFolderDataCollection `xml:"m:EncryptedSharedFolderDataCollection"`
	// The InvalidRecipients element represents the recipients of a folder sharing request that are invalid.
	InvalidRecipients *InvalidRecipients `xml:"m:InvalidRecipients"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Describes the status of the response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	GetSharingMetadataResponseSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - The Active Directory directory service is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	GetSharingMetadataResponseWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempt by any client  - Server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	GetSharingMetadataResponseError = `Error`
)
