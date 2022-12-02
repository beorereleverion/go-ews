package elements

// The CopyFolderResponse element defines a response to a CopyFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyfolderresponse
type CopyFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
