package elements

// The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responseobjects
type ResponseObjects struct {
	// The AcceptItem element represents an Accept reply to a meeting request.
	AcceptItem *AcceptItem `xml:"t:AcceptItem"`
	// The AcceptSharingInvitation element is used to accept an invitation that allows access to another user's calendar or contacts data.
	AcceptSharingInvitation *AcceptSharingInvitation `xml:"t:AcceptSharingInvitation"`
	// The CancelCalendarItem element represents the response object that is used to cancel a meeting.
	CancelCalendarItem *CancelCalendarItem `xml:"t:CancelCalendarItem"`
	// The DeclineItem element represents a Decline reply to a meeting request.
	DeclineItem *DeclineItem `xml:"t:DeclineItem"`
	// The ForwardItem element contains an Exchange store item to forward to recipients.
	ForwardItem *ForwardItem `xml:"t:ForwardItem"`
	// The PostReplyItem element contains a reply to a post item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	PostReplyItem *PostReplyItem `xml:"t:PostReplyItem"`
	// The ProposeNewTime element specifies a response object that indicates that the meeting attendee can propose a new meeting time.
	ProposeNewTime *ProposeNewTime `xml:"t:ProposeNewTime"`
	// The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
	RemoveItem *RemoveItem `xml:"t:RemoveItem"`
	// The ReplyToAllItem element contains a reply to the sender and all identified recipients of an item in the Exchange store.
	ReplyAllToItem *ReplyAllToItem `xml:"t:ReplyAllToItem"`
	// The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
	ReplyToItem *ReplyToItem `xml:"t:ReplyToItem"`
	// The SuppressReadReceipt element is used to suppress read receipts.
	SuppressReadReceipt *SuppressReadReceipt `xml:"t:SuppressReadReceipt"`
	// The TentativelyAcceptItem element represents a Tentative reply to a meeting request.
	TentativelyAcceptItem *TentativelyAcceptItem `xml:"t:TentativelyAcceptItem"`
}
