package elements

// The ExpandDLResponseMessage element contains the status and result of a single ExpandDL operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expanddlresponsemessage
type ExpandDLResponseMessage struct {
	// The DLExpansion element contains an array of mailboxes that are contained in a distribution list.
	DLExpansion *DLExpansion `xml:"t:DLExpansion"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Represents the next denominator to use for the next request when doing fractional paging.
	AbsoluteDenominator *string `xml:"AbsoluteDenominator,attr"`
	// Indicates that additional paging is not needed. This attribute will be true if the current results contain the last item in the query.
	IncludesLastItemInRange *string `xml:"IncludesLastItemInRange,attr"`
	// Represents the next index that should be used for the next request when an indexed paging view is used.
	IndexedPagingOffset *string `xml:"IndexedPagingOffset,attr"`
	// Represents the new numerator value to use for the next request when fraction page views are used.
	NumeratorOffset *string `xml:"NumeratorOffset,attr"`
	// Describes the status of an ExpandDL operation response.The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
	// Represents the total number of items that pass the restriction.
	TotalItemsInView *string `xml:"TotalItemsInView,attr"`
}

const (
	// Describes a request that is fulfilled.
	ExpandDLResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes are moved.  - The mailbox database (MDB) is offline.  - A password is expired.  - A quota was exceeded.
	ExpandDLResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements  - Attributes or elements that are out of range  - An unknown tag  - An attribute or element that is not valid in the context  - An unauthorized access attempt by any client  - A server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	ExpandDLResponseMessageError = `Error`
)
