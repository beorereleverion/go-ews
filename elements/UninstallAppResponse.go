package elements

// The UninstallAppResponse element specifies the response to an UninstallApp request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uninstallappresponse
type UninstallAppResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
