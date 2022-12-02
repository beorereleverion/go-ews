package elements

// The DeleteUserConfigurationResponse element defines a response to a single DeleteUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteuserconfigurationresponse
type DeleteUserConfigurationResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
