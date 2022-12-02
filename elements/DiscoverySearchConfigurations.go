package elements

// The DiscoverySearchConfigurations element specifies an array of DiscoverySearchConfiguration elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/discoverysearchconfigurations
type DiscoverySearchConfigurations struct {
	// The DiscoverySearchConfiguration element specifies the configuration for eDiscovery search.
	DiscoverySearchConfiguration *DiscoverySearchConfiguration `xml:"m:DiscoverySearchConfiguration"`
}
