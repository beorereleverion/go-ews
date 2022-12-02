package elements

// The MeetingResponse element represents a meeting response in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingresponse
type MeetingResponse struct {
	// The AssociatedCalendarItemId element represents the calendar item that is associated with a MeetingMessage, MeetingRequest, MeetingResponse, MeetingCancellation, or ReminderMessageData.
	AssociatedCalendarItemId *AssociatedCalendarItemId `xml:"t:AssociatedCalendarItemId"`
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"t:Attachments"`
	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"t:BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"t:Body"`
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"t:Categories"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"t:CcRecipients"`
	// The ConversationIndex element contains a binary ID that represents the thread to which this message belongs.
	ConversationIndex *ConversationIndex `xml:"t:ConversationIndex"`
	// The ConversationTopic element represents the conversation topic.
	ConversationTopic *ConversationTopic `xml:"t:ConversationTopic"`
	// The Culture element represents the culture for a given item in a mailbox.
	Culture *Culture `xml:"t:Culture"`
	// The DateTimeCreated element represents the date and time that an item in the mailbox was created.
	DateTimeCreated *DateTimeCreated `xml:"t:DateTimeCreated"`
	// The DateTimeReceived element represents the date and time that an item in a mailbox was received.
	DateTimeReceived *DateTimeReceived `xml:"t:DateTimeReceived"`
	// The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
	DateTimeSent *DateTimeSent `xml:"t:DateTimeSent"`
	// The DateTimeStamp element indicates the date and time that an instance of a calendar object was created.
	DateTimeStamp *DateTimeStamp `xml:"t:DateTimeStamp"`
	// The DisplayCc element represents the display string that is used for the contents of the Cc box. This is the concatenated string of all Cc recipient display names.
	DisplayCc *DisplayCc `xml:"t:DisplayCc"`
	// The DisplayTo element represents the display string that is used for the contents of the To box. This is the concatenated string of all To recipient display names.
	DisplayTo *DisplayTo `xml:"t:DisplayTo"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty"`
	// The From element represents the address from which the message was sent.
	From *From `xml:"t:From"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"t:HasAttachments"`
	// The HasBeenProcessed element indicates whether a meeting message item has been processed.
	HasBeenProcessed *HasBeenProcessed `xml:"t:HasBeenProcessed"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"t:Importance"`
	// The InReplyTo element represents the identifier of the item to which this item is a reply.
	InReplyTo *InReplyTo `xml:"t:InReplyTo"`
	// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	InternetMessageHeaders *InternetMessageHeaders `xml:"t:InternetMessageHeaders"`
	// The InternetMessageId element represents the Internet message identifier of an item.
	InternetMessageId *InternetMessageId `xml:"t:InternetMessageId"`
	// The IsDelegated element indicates whether a meeting was handled by an account that has delegate access.
	IsDelegated *IsDelegated `xml:"t:IsDelegated"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"t:IsDeliveryReceiptRequested"`
	// The IsDraft element indicates whether an item has not yet been sent.
	IsDraft *IsDraft `xml:"t:IsDraft"`
	// The IsFromMe element indicates whether a user sent an item to him or herself.
	IsFromMe *IsFromMe `xml:"t:IsFromMe"`
	// The IsOutOfDate element indicates whether a meeting message, request, response, or cancellation is out-of-date.
	IsOutOfDate *IsOutOfDate `xml:"t:IsOutOfDate"`
	// The IsRead element indicates whether a message has been read.
	IsRead *IsRead `xml:"t:IsRead"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"t:IsReadReceiptRequested"`
	// The IsResend element indicates whether the item had previously been sent.
	IsResend *IsResend `xml:"t:IsResend"`
	// The IsResponseRequested element indicates whether a response to an item is requested.
	IsResponseRequested *IsResponseRequested `xml:"t:IsResponseRequested"`
	// The IsSubmitted element indicates whether an item has been submitted to the Outbox default folder.
	IsSubmitted *IsSubmitted `xml:"t:IsSubmitted"`
	// The IsUnmodified element indicates whether the item has been modified.
	IsUnmodified *IsUnmodified `xml:"t:IsUnmodified"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"t:ItemClass"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"t:MeetingMessage"`
	// The MimeContent element contains the ASCII MIME stream of an object that is represented in base64Binary format and supports [RFC2045].
	MimeContent *MimeContent `xml:"t:MimeContent"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// The ProposedEnd element specifies the proposed end time of a meeting.
	ProposedEnd *ProposedEnd `xml:"t:ProposedEnd"`
	// The ProposedStart element specifies the proposed start time of a meeting.
	ProposedStart *ProposedStart `xml:"t:ProposedStart"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"t:ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"t:ReceivedRepresenting"`
	// The RecurrenceId element is used to identify a specific instance of a recurring calendar item.
	RecurrenceId *RecurrenceId `xml:"t:RecurrenceId"`
	// The References element represents the Usenet header that is used to associate replies with the original messages.
	References *References `xml:"t:References"`
	// The ReminderDueBy element represents the date and time when the event occurs. This is used by the ReminderMinutesBeforeStart element to determine when the reminder is displayed.
	ReminderDueBy *ReminderDueBy `xml:"t:ReminderDueBy"`
	// The ReminderIsSet element indicates whether a reminder has been set for an item in the Exchange store.
	ReminderIsSet *ReminderIsSet `xml:"t:ReminderIsSet"`
	// The ReminderMinutesBeforeStart element represents the number of minutes before an event occurs when a reminder is displayed.
	ReminderMinutesBeforeStart *ReminderMinutesBeforeStart `xml:"t:ReminderMinutesBeforeStart"`
	// The ReplyTo element identifies an array of addresses to which replies should be sent.
	ReplyTo *ReplyTo `xml:"t:ReplyTo"`
	// The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
	ResponseObjects *ResponseObjects `xml:"t:ResponseObjects"`
	// The ResponseType element represents the type of recipient response that is received for a meeting.
	ResponseType *ResponseType `xml:"t:ResponseType"`
	// The Sender element identifies the sender of an item.
	Sender *Sender `xml:"t:Sender"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"t:Sensitivity"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"t:Size"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"t:Subject"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients *ToRecipients `xml:"t:ToRecipients"`
	// The UID element uniquely identifies a calendar item.
	UID *UID `xml:"t:UID"`
}
