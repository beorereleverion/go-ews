package elements

// The GetDiscoverySearchConfigurationResponse element specifies the response to a GetDiscoverySearchConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdiscoverysearchconfigurationresponse
import "encoding/xml"

type GetDiscoverySearchConfigurationResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The DiscoverySearchConfigurations element specifies an array of DiscoverySearchConfiguration elements.
	DiscoverySearchConfigurations *DiscoverySearchConfigurations `xml:"DiscoverySearchConfigurations"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (G *GetDiscoverySearchConfigurationResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetDiscoverySearchConfigurationResponse"
}

func (G *GetDiscoverySearchConfigurationResponse) GetSchema() *Schema {
	return &SchemaMessages
}
