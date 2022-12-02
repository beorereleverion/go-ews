package elements

// The DeleteFolderResponse element defines a response to a DeleteFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deletefolderresponse
type DeleteFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
