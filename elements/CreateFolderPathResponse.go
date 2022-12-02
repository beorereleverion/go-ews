package elements

// The CreateFolderPathResponse element is used to return a folder path.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderpathresponse
type CreateFolderPathResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
