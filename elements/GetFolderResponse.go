package elements

// The GetFolderResponse element defines a response to a GetFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponse
type GetFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
