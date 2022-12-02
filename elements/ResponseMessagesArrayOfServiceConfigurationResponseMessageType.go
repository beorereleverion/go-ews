package elements

// The ResponseMessages element contains an array of service configuration response messages.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages-arrayofserviceconfigurationresponsemessagetype
type ResponseMessagesArrayOfServiceConfigurationResponseMessageType struct {
	// The ServiceConfigurationResponseMessageType element contains service configuration settings.
	ServiceConfigurationResponseMessageType *ServiceConfigurationResponseMessageType `xml:"m:ServiceConfigurationResponseMessageType"`
}
