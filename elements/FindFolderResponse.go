package elements

// The FindFolderResponse element defines a response to a FindFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findfolderresponse
type FindFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
