package elements

// The DeliverMeetingRequests element defines how meeting requests are handled between the delegate and the principal. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delivermeetingrequests
type DeliverMeetingRequests string

const (
	// DelegatesAndMe
	DeliverMeetingRequestsDelegatesAndMe DeliverMeetingRequests = `DelegatesAndMe`
	// DelegatesAndSendInformationToMe
	DeliverMeetingRequestsDelegatesAndSendInformationToMe DeliverMeetingRequests = `DelegatesAndSendInformationToMe`
	// DelegatesOnly
	DeliverMeetingRequestsDelegatesOnly DeliverMeetingRequests = `DelegatesOnly`
	// NoForward
	DeliverMeetingRequestsNoForward DeliverMeetingRequests = `NoForward`
)
