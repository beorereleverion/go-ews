package elements

// The GetAppMarketplaceUrlResponse element specifies the response to a GetAppMarketplaceUrl request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmarketplaceurlresponse
type GetAppMarketplaceUrlResponse struct {
	// The AppMarketplaceUrl element specifies the URL for the app marketplace.
	AppMarketplaceUrl *AppMarketplaceUrl `xml:"m:AppMarketplaceUrl"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"m:MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"m:ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}
