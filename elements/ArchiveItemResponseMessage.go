package elements

// The ArchiveItemResponseMessage element specifies the response message to an ArchiveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archiveitemresponsemessage
type ArchiveItemResponseMessage struct {
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
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
