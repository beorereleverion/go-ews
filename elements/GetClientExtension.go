package elements

// The GetClientExtension element represents a request to get a client extension.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getclientextension
type GetClientExtension struct {
	// The IsDebug element is not used.
	IsDebug *IsDebug `xml:"m:IsDebug"`
	// The RequestedExtensionIds element contains an array of extension identifiers.
	RequestedExtensionIds *RequestedExtensionIds `xml:"m:RequestedExtensionIds"`
	// The UserParameters element contains a list of enabled and disabled client extensions.
	UserParameters *UserParameters `xml:"t:UserParameters"`
}
