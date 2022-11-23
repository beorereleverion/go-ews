package elements

// The ResponseMessages element contains the response messages for an Exchange Web Services request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsemessages
type ResponseMessages struct {
	GetFolderResponseMessage string `xml:"m:GetFolderResponseMessage"`
}
