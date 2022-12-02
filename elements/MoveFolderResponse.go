package elements

// The MoveFolderResponse element defines a response to a MoveFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/movefolderresponse
type MoveFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
