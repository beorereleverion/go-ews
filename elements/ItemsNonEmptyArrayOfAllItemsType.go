package elements

// The Items element contains a set of items to create.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/items-nonemptyarrayofallitemstype
type ItemsNonEmptyArrayOfAllItemsType struct {
	// The AcceptItem element represents an Accept reply to a meeting request.
	AcceptItem *AcceptItem `xml:"t:AcceptItem"`
	// The AcceptSharingInvitation element is used to accept an invitation that allows access to another user's calendar or contacts data.
	AcceptSharingInvitation *AcceptSharingInvitation `xml:"t:AcceptSharingInvitation"`
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
	// The CancelCalendarItem element represents the response object that is used to cancel a meeting.
	CancelCalendarItem *CancelCalendarItem `xml:"t:CancelCalendarItem"`
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"t:Contact"`
	// The DeclineItem element represents a Decline reply to a meeting request.
	DeclineItem *DeclineItem `xml:"t:DeclineItem"`
	// The DistributionList element represents a distribution list.
	DistributionList *DistributionList `xml:"t:DistributionList"`
	// The ForwardItem element contains an Exchange store item to forward to recipients.
	ForwardItem *ForwardItem `xml:"t:ForwardItem"`
	// The Item element represents a generic item in the Exchange store.
	Item *Item `xml:"t:Item"`
	// The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	MeetingCancellation *MeetingCancellation `xml:"t:MeetingCancellation"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"t:MeetingMessage"`
	// The MeetingRequest element represents a meeting request in the Exchange store.
	MeetingRequest *MeetingRequest `xml:"t:MeetingRequest"`
	// The MeetingResponse element represents a meeting response in the Exchange store.
	MeetingResponse *MeetingResponse `xml:"t:MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *Message `xml:"t:Message"`
	// The PostReplyItem element contains a reply to a post item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	PostReplyItem *PostReplyItem `xml:"t:PostReplyItem"`
	// The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
	RemoveItem *RemoveItem `xml:"t:RemoveItem"`
	// The ReplyToAllItem element contains a reply to the sender and all identified recipients of an item in the Exchange store.
	ReplyAllToItem *ReplyAllToItem `xml:"t:ReplyAllToItem"`
	// The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
	ReplyToItem *ReplyToItem `xml:"t:ReplyToItem"`
	// The SuppressReadReceipt element is used to suppress read receipts.
	SuppressReadReceipt *SuppressReadReceipt `xml:"t:SuppressReadReceipt"`
	// The Task element represents a task in the Exchange store.
	Task *Task `xml:"t:Task"`
	// The TentativelyAcceptItem element represents a Tentative reply to a meeting request.
	TentativelyAcceptItem *TentativelyAcceptItem `xml:"t:TentativelyAcceptItem"`
}
