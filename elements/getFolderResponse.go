package elements

// The GetFolderResponse element defines a response to a GetFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponse
type GetFolderResponse struct {
	// Contains the status and result of a single GetFolder request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}
