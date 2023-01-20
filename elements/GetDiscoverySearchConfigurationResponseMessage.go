package elements

// The GetDiscoverySearchConfigurationResponseMessage element specifies the response message for a GetDiscoverySearchConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getdiscoverysearchconfigurationresponsemessage
import "encoding/xml"

type GetDiscoverySearchConfigurationResponseMessage struct {
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

func (G *GetDiscoverySearchConfigurationResponseMessage) SetForMarshal() {
	G.XMLName.Local = "m:GetDiscoverySearchConfigurationResponseMessage"
}

func (G *GetDiscoverySearchConfigurationResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
