package elements

// The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responseobjects
import "encoding/xml"

type ResponseObjects struct {
	XMLName xml.Name

	// The AcceptItem element represents an Accept reply to a meeting request.
	AcceptItem *AcceptItem `xml:"AcceptItem"`
	// The AcceptSharingInvitation element is used to accept an invitation that allows access to another user's calendar or contacts data.
	AcceptSharingInvitation *AcceptSharingInvitation `xml:"AcceptSharingInvitation"`
	// The CancelCalendarItem element represents the response object that is used to cancel a meeting.
	CancelCalendarItem *CancelCalendarItem `xml:"CancelCalendarItem"`
	// The DeclineItem element represents a Decline reply to a meeting request.
	DeclineItem *DeclineItem `xml:"DeclineItem"`
	// The ForwardItem element contains an Exchange store item to forward to recipients.
	ForwardItem *ForwardItem `xml:"ForwardItem"`
	// The PostReplyItem element contains a reply to a post item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	PostReplyItem *PostReplyItem `xml:"PostReplyItem"`
	// The ProposeNewTime element specifies a response object that indicates that the meeting attendee can propose a new meeting time.
	ProposeNewTime *ProposeNewTime `xml:"ProposeNewTime"`
	// The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
	RemoveItem *RemoveItem `xml:"RemoveItem"`
	// The ReplyToAllItem element contains a reply to the sender and all identified recipients of an item in the Exchange store.
	ReplyAllToItem *ReplyAllToItem `xml:"ReplyAllToItem"`
	// The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
	ReplyToItem *ReplyToItem `xml:"ReplyToItem"`
	// The SuppressReadReceipt element is used to suppress read receipts.
	SuppressReadReceipt *SuppressReadReceipt `xml:"SuppressReadReceipt"`
	// The TentativelyAcceptItem element represents a Tentative reply to a meeting request.
	TentativelyAcceptItem *TentativelyAcceptItem `xml:"TentativelyAcceptItem"`
}

func (R *ResponseObjects) SetForMarshal() {
	R.XMLName.Local = "t:ResponseObjects"
}

func (R *ResponseObjects) GetSchema() *Schema {
	return &SchemaTypes
}
