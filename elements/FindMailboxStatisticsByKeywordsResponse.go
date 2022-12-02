package elements

// The FindMailboxStatisticsByKeywordsResponse element specifies the response to a FindMailboxStatisticsByKeywords request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmailboxstatisticsbykeywordsresponse
type FindMailboxStatisticsByKeywordsResponse struct {
	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"m:ResponseMessages"`
}
