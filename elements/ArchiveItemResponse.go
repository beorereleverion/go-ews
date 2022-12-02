package elements

// The ArchiveItemResponse element specifies the response to an ArchiveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archiveitemresponse
type ArchiveItemResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
