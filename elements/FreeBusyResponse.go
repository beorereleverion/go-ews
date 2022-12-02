package elements

// The FreeBusyResponse element contains the free/busy information for a single mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponse
type FreeBusyResponse struct {
	// The FreeBusyView element contains availability information for a specific user.
	FreeBusyView *FreeBusyView `xml:"t:FreeBusyView"`
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"m:ResponseMessage"`
}
