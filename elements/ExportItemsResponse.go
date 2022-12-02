package elements

// The ExportItemsResponse element represents response to a single ExportItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exportitemsresponse
type ExportItemsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
