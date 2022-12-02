package elements

// The UploadItemsResponse element represents a response to a single UploadItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uploaditemsresponse
type UploadItemsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
