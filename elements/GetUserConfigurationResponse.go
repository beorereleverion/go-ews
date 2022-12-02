package elements

// The GetUserConfigurationResponse element defines a response to a single GetUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserconfigurationresponse
type GetUserConfigurationResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
