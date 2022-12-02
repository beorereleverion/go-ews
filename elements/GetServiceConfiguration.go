package elements

// The GetServiceConfiguration element defines a GetServiceConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getserviceconfiguration
type GetServiceConfiguration struct {
	// The ActingAs element identifies who the caller is sending as.
	ActingAs *ActingAs `xml:"m:ActingAs"`
	// The RequestedConfiguration element contains the requested service configurations.
	RequestedConfiguration *RequestedConfiguration `xml:"m:RequestedConfiguration"`
}
