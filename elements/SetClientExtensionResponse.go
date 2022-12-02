package elements

// The SetClientExtensionResponse element contains the response to a SetClientExtension request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setclientextensionresponse
type SetClientExtensionResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
