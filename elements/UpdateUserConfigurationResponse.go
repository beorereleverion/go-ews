package elements

// The UpdateUserConfigurationResponse element defines a response to a single UpdateUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateuserconfigurationresponse
type UpdateUserConfigurationResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
