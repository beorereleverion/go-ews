package elements

// The CreateManagedFolderResponse element defines a response to a CreateManagedFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createmanagedfolderresponse
type CreateManagedFolderResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
