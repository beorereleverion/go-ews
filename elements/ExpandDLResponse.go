package elements

// The ExpandDLResponse element defines a response to a request to expand a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expanddlresponse
type ExpandDLResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
