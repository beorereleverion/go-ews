package elements

// The RequestedConfiguration element contains the requested service configurations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/requestedconfiguration
type RequestedConfiguration struct {
	// The ConfigurationName element specifies the requested service configurations by name.
	ConfigurationName *ConfigurationName `xml:"m:ConfigurationName"`
}
