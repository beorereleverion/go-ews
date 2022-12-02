package elements

// The UpdateItemInRecoverableItemsResponse element specifies the response to an UpdateItemInRecoverableItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateiteminrecoverableitemsresponse
type UpdateItemInRecoverableItemsResponse struct {
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"t:Attachments"`
	// The ConflictResults element contains the number of conflicts in an UpdateItem operation response.
	ConflictResults *ConflictResults `xml:"t:ConflictResults"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The Items element contains an array of items.
	Items *Items `xml:"t:Items"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
