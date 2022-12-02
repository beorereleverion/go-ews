package elements

// The SearchMailboxesResponse element contains the response to a SearchMailboxes WSDL operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchmailboxesresponse
type SearchMailboxesResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
