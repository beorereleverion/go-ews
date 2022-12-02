package elements

// The GetDiscoverySearchConfigurationResponse element specifies the response to a GetDiscoverySearchConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdiscoverysearchconfigurationresponse
type GetDiscoverySearchConfigurationResponse struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The DiscoverySearchConfigurations element specifies an array of DiscoverySearchConfiguration elements.
	DiscoverySearchConfigurations *DiscoverySearchConfigurations `xml:"m:DiscoverySearchConfigurations"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
}
