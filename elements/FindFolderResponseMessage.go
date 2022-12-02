package elements

// The FindFolderResponseMessage element contains the status and result of a single FindFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolderresponsemessage
type FindFolderResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// The RootFolder element contains the results of a search of a single root folder during a FindFolder operation.
	RootFolder *RootFolderFindFolderResponseMessage `xml:"m:RootFolder"`
	// Describes the status of a FindFolder operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}
