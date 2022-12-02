package elements

// The CreateFolderResponse element defines a response to a CreateFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderresponse
type CreateFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
