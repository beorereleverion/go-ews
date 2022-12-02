package elements

// The UpdateFolderResponse element defines the response to an UpdateFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolderresponse
type UpdateFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
