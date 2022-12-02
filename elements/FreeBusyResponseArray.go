package elements

// The FreeBusyResponseArray element contains the requested users' availability information and the response status.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponsearray
type FreeBusyResponseArray struct {
	// The FreeBusyResponse element contains the free/busy information for a single mailbox user.
	FreeBusyResponse *FreeBusyResponse `xml:"m:FreeBusyResponse"`
}
