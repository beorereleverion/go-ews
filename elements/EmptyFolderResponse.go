package elements

// The EmptyFolderResponse element defines a response to an EmptyFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emptyfolderresponse
type EmptyFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
