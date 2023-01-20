package elements

// The GetAppMarketplaceUrlResponse element specifies the response to a GetAppMarketplaceUrl request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getappmarketplaceurlresponse
import "encoding/xml"

type GetAppMarketplaceUrlResponse struct {
	XMLName xml.Name

	// The AppMarketplaceUrl element specifies the URL for the app marketplace.
	AppMarketplaceUrl *AppMarketplaceUrl `xml:"AppMarketplaceUrl"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (G *GetAppMarketplaceUrlResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetAppMarketplaceUrlResponse"
}

func (G *GetAppMarketplaceUrlResponse) GetSchema() *Schema {
	return &SchemaMessages
}
