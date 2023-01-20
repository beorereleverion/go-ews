package elements

// The DiscoverySearchConfigurations element specifies an array of DiscoverySearchConfiguration elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/discoverysearchconfigurations
import "encoding/xml"

type DiscoverySearchConfigurations struct {
	XMLName xml.Name

	// The DiscoverySearchConfiguration element specifies the configuration for eDiscovery search.
	DiscoverySearchConfiguration *DiscoverySearchConfiguration `xml:"DiscoverySearchConfiguration"`
}

func (D *DiscoverySearchConfigurations) SetForMarshal() {
	D.XMLName.Local = "m:DiscoverySearchConfigurations"
}

func (D *DiscoverySearchConfigurations) GetSchema() *Schema {
	return &SchemaMessages
}
