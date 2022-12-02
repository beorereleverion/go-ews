package elements

// The SyncFolderItemsResponse element defines a response to a SyncFolderItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderitemsresponse
type SyncFolderItemsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
