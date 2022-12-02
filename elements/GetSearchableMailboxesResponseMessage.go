package elements

// The GetSearchableMailboxesResponseMessage element specifies the response message for a GetSearchableMailboxes request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getsearchablemailboxesresponsemessage
type GetSearchableMailboxesResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The SearchableMailboxes element contains an array of the mailboxes returned from a GetSearchableMailboxes request.
	SearchableMailboxes *SearchableMailboxes `xml:"m:SearchableMailboxes"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
