package elements

// The CreateUserConfigurationResponse element defines a response to a single CreateUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createuserconfigurationresponse
type CreateUserConfigurationResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
