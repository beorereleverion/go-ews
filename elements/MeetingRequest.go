package elements

// The MeetingRequest element represents a meeting request in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingrequest
import "encoding/xml"

type MeetingRequest struct {
	XMLName xml.Name

	// The AdjacentMeetingCount element represents the total number of calendar items that are adjacent to a meeting time.
	AdjacentMeetingCount *AdjacentMeetingCount `xml:"AdjacentMeetingCount"`
	// The AdjacentMeetings element identifies all calendar items that are adjacent to a meeting time.
	AdjacentMeetings *AdjacentMeetings `xml:"AdjacentMeetings"`
	// The AllowNewTimeProposal element indicates whether a new meeting time can be proposed for a meeting by an attendee.
	AllowNewTimeProposal *AllowNewTimeProposal `xml:"AllowNewTimeProposal"`
	// The AppointmentReplyTime element represents the date and time that an attendee replied to a meeting request.
	AppointmentReplyTime *AppointmentReplyTime `xml:"AppointmentReplyTime"`
	// The AppointmentSequenceNumber element specifies the sequence number of a version of an appointment.
	AppointmentSequenceNumber *AppointmentSequenceNumber `xml:"AppointmentSequenceNumber"`
	// The AppointmentState element specifies the status of the appointment.
	AppointmentState *AppointmentState `xml:"AppointmentState"`
	// The AssociatedCalendarItemId element represents the calendar item that is associated with a MeetingMessage, MeetingRequest, MeetingResponse, MeetingCancellation, or ReminderMessageData.
	AssociatedCalendarItemId *AssociatedCalendarItemId `xml:"AssociatedCalendarItemId"`
	// The Attachments element contains the items or files that are attached to an item in the Exchange store.
	Attachments *Attachments `xml:"Attachments"`
	// The BccRecipients element represents a collection of recipients to receive a blind carbon copy (Bcc) of an e-mail message.
	BccRecipients *BccRecipients `xml:"BccRecipients"`
	// The Body element specifies the body of an item.
	Body *Body `xml:"Body"`
	// The CalendarItemType element represents the type of a calendar item.
	CalendarItemType *CalendarItemType `xml:"CalendarItemType"`
	// The Categories element contains a collection of strings that identify the categories to which an item in the mailbox belongs.
	Categories *Categories `xml:"Categories"`
	// The CcRecipients element represents a collection of recipients that will receive a copy of the message.
	CcRecipients *CcRecipients `xml:"CcRecipients"`
	// The ConferenceType element describes the type of conferencing that is performed with a calendar item.
	ConferenceType *ConferenceType `xml:"ConferenceType"`
	// The ConflictingMeetingCount element represents the number of meetings that conflict with the calendar item.
	ConflictingMeetingCount *ConflictingMeetingCount `xml:"ConflictingMeetingCount"`
	// The ConflictingMeetings element identifies all calendar items that conflict with a meeting time.
	ConflictingMeetings *ConflictingMeetings `xml:"ConflictingMeetings"`
	// The ConversationId element contains the identifier of an item or conversation.
	ConversationId *ConversationId `xml:"ConversationId"`
	// The ConversationIndex element contains a binary ID that represents the thread to which this message belongs.
	ConversationIndex *ConversationIndex `xml:"ConversationIndex"`
	// The ConversationTopic element represents the conversation topic.
	ConversationTopic *ConversationTopic `xml:"ConversationTopic"`
	// The Culture element represents the culture for a given item in a mailbox.
	Culture *Culture `xml:"Culture"`
	// The DateTimeCreated element represents the date and time that an item in the mailbox was created.
	DateTimeCreated *DateTimeCreated `xml:"DateTimeCreated"`
	// The DateTimeReceived element represents the date and time that an item in a mailbox was received.
	DateTimeReceived *DateTimeReceived `xml:"DateTimeReceived"`
	// The DateTimeSent element represents the date and time at which an item in a mailbox was sent.
	DateTimeSent *DateTimeSent `xml:"DateTimeSent"`
	// The DateTimeStamp element indicates the date and time that an instance of a calendar object was created.
	DateTimeStamp *DateTimeStamp `xml:"DateTimeStamp"`
	// The DeletedOccurrences element contains an array of deleted occurrences of a recurring calendar item.
	DeletedOccurrences *DeletedOccurrences `xml:"DeletedOccurrences"`
	// The DisplayCc element represents the display string that is used for the contents of the Cc box. This is the concatenated string of all Cc recipient display names.
	DisplayCc *DisplayCc `xml:"DisplayCc"`
	// The DisplayTo element represents the display string that is used for the contents of the To box. This is the concatenated string of all To recipient display names.
	DisplayTo *DisplayTo `xml:"DisplayTo"`
	// The Duration element represents the duration of a calendar item.
	Duration *DurationItems `xml:"Duration"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"EffectiveRights"`
	// The End element represents the end of a duration.
	End *End `xml:"End"`
	// The EndTimeZone element defines the time zone for the end time of a CalendarItem or MeetingRequest.
	EndTimeZone *EndTimeZone `xml:"EndTimeZone"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"ExtendedProperty"`
	// The FirstOccurrence element represents the first occurrence of a recurring calendar item.
	FirstOccurrence *FirstOccurrence `xml:"FirstOccurrence"`
	// The From element represents the address from which the message was sent.
	From *From `xml:"From"`
	// The HasAttachments element represents a property that is set to true if an item has at least one visible attachment or if a conversation contains at least one item that has an attachment. This property is read-only.
	HasAttachments *HasAttachments `xml:"HasAttachments"`
	// The HasBeenProcessed element indicates whether a meeting message item has been processed.
	HasBeenProcessed *HasBeenProcessed `xml:"HasBeenProcessed"`
	// The Importance element describes the importance of an item or the aggregated importance of all items in a conversation in the current folder.
	Importance *Importance `xml:"Importance"`
	// The InReplyTo element represents the identifier of the item to which this item is a reply.
	InReplyTo *InReplyTo `xml:"InReplyTo"`
	// The IntendedFreeBusyStatus element represents the intended status for the calendar item that is associated with the meeting request.
	IntendedFreeBusyStatus *IntendedFreeBusyStatus `xml:"IntendedFreeBusyStatus"`
	// The InternetMessageHeaders element contains a collection of some of the Internet message headers that are contained in an item in a mailbox. To get the entire collection of Internet message headers, use the PR_TRANSPORT_MESSAGE_HEADERS property. For more information about EWS and Internet message headers, seeGetting Internet message headersin EWS, MIME, and the missing Internet message headers.
	InternetMessageHeaders *InternetMessageHeaders `xml:"InternetMessageHeaders"`
	// The InternetMessageId element represents the Internet message identifier of an item.
	InternetMessageId *InternetMessageId `xml:"InternetMessageId"`
	// The IsAllDayEvent element indicates whether a calendar item or meeting request represents an all-day event.
	IsAllDayEvent *IsAllDayEvent `xml:"IsAllDayEvent"`
	// The IsAssociated element indicates whether the item is associated with a folder.
	IsAssociated *IsAssociated `xml:"IsAssociated"`
	// The IsCancelled element indicates whether an appointment or meeting has been canceled.
	IsCancelled *IsCancelled `xml:"IsCancelled"`
	// The IsDelegated element indicates whether a meeting was handled by an account that has delegate access.
	IsDelegated *IsDelegated `xml:"IsDelegated"`
	// The IsDeliveryReceiptRequested element indicates whether the sender of an item requests a delivery receipt.
	IsDeliveryReceiptRequested *IsDeliveryReceiptRequested `xml:"IsDeliveryReceiptRequested"`
	// The IsDraft element indicates whether an item has not yet been sent.
	IsDraft *IsDraft `xml:"IsDraft"`
	// The IsFromMe element indicates whether a user sent an item to him or herself.
	IsFromMe *IsFromMe `xml:"IsFromMe"`
	// The IsMeeting element indicates whether the calendar item is a meeting or an appointment.
	IsMeeting *IsMeeting `xml:"IsMeeting"`
	// The IsOnlineMeeting element indicates whether the meeting is online.
	IsOnlineMeeting *IsOnlineMeeting `xml:"IsOnlineMeeting"`
	// The IsOutOfDate element indicates whether a meeting message, request, response, or cancellation is out-of-date.
	IsOutOfDate *IsOutOfDate `xml:"IsOutOfDate"`
	// The IsRead element indicates whether a message has been read.
	IsRead *IsRead `xml:"IsRead"`
	// The IsReadReceiptRequested element indicates whether the sender of an item requests a read receipt.
	IsReadReceiptRequested *IsReadReceiptRequested `xml:"IsReadReceiptRequested"`
	// The IsRecurring element indicates whether a calendar item, meeting request, or task is part of a recurring item. This element is read-only.
	IsRecurring *IsRecurring `xml:"IsRecurring"`
	// The IsResend element indicates whether the item had previously been sent.
	IsResend *IsResend `xml:"IsResend"`
	// The IsResponseRequested element indicates whether a response to an item is requested.
	IsResponseRequested *IsResponseRequested `xml:"IsResponseRequested"`
	// The IsSubmitted element indicates whether an item has been submitted to the Outbox default folder.
	IsSubmitted *IsSubmitted `xml:"IsSubmitted"`
	// The IsUnmodified element indicates whether the item has been modified.
	IsUnmodified *IsUnmodified `xml:"IsUnmodified"`
	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
	// The LastModifiedName element contains the display name of the last user to modify an item. This element is read-only. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	LastModifiedName *LastModifiedName `xml:"LastModifiedName"`
	// The LastModifiedTime element indicates when an item was last modified. This element is read-only.
	LastModifiedTime *LastModifiedTime `xml:"LastModifiedTime"`
	// The LastOccurrence element represents the last occurrence of a recurring calendar item.
	LastOccurrence *LastOccurrence `xml:"LastOccurrence"`
	// The LegacyFreeBusyStatus element represents the free/busy status of the calendar item.
	LegacyFreeBusyStatus *LegacyFreeBusyStatus `xml:"LegacyFreeBusyStatus"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"Location"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"MeetingMessage"`
	// The MeetingRequestType element describes the type of the meeting request.
	MeetingRequestType *MeetingRequestType `xml:"MeetingRequestType"`
	// The MeetingRequestWasSent element indicates whether a meeting request has been sent to requested attendees.
	MeetingRequestWasSent *MeetingRequestWasSent `xml:"MeetingRequestWasSent"`
	// The MeetingTimeZone element represents the time zone of the location where the meeting is hosted.
	MeetingTimeZone *MeetingTimeZone `xml:"MeetingTimeZone"`
	// The MeetingWorkspaceUrl element contains the URL for the meeting workspace that is included in the calendar item. A meeting workspace is a shared Web site for planning the meeting and tracking the results.
	MeetingWorkspaceUrl *MeetingWorkspaceUrl `xml:"MeetingWorkspaceUrl"`
	// The MimeContent element contains the ASCII MIME stream of an object that is represented in base64Binary format and supports [RFC2045].
	MimeContent *MimeContent `xml:"MimeContent"`
	// The ModifiedOccurrences element contains an array of recurring calendar item occurrences that have been modified so that they are different than the recurrence master item.
	ModifiedOccurrences *ModifiedOccurrences `xml:"ModifiedOccurrences"`
	// The MyResponseType element contains the status of or response to a calendar item.
	MyResponseType *MyResponseType `xml:"MyResponseType"`
	// The NetShowUrl element specifies the URL for a Microsoft NetShow online meeting.
	NetShowUrl *NetShowUrl `xml:"NetShowUrl"`
	// The OptionalAttendees element represents attendees who are not required to attend a meeting.
	OptionalAttendees *OptionalAttendees `xml:"OptionalAttendees"`
	// The Organizer element represents the organizer of a meeting.
	Organizer *Organizer `xml:"Organizer"`
	// The OriginalStart element represents the original start time of a calendar item.
	OriginalStart *OriginalStart `xml:"OriginalStart"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The ReceivedBy element identifies the delegate in a delegate access scenario.
	ReceivedBy *ReceivedBy `xml:"ReceivedBy"`
	// The ReceivedRepresenting element identifies the principal in a delegate access scenario.
	ReceivedRepresenting *ReceivedRepresenting `xml:"ReceivedRepresenting"`
	// The Recurrence element contains the recurrence pattern for calendar items and meeting requests.
	Recurrence *RecurrenceRecurrenceType `xml:"Recurrence"`
	// The RecurrenceId element is used to identify a specific instance of a recurring calendar item.
	RecurrenceId *RecurrenceId `xml:"RecurrenceId"`
	// The References element represents the Usenet header that is used to associate replies with the original messages.
	References *References `xml:"References"`
	// The ReminderDueBy element represents the date and time when the event occurs. This is used by the ReminderMinutesBeforeStart element to determine when the reminder is displayed.
	ReminderDueBy *ReminderDueBy `xml:"ReminderDueBy"`
	// The ReminderIsSet element indicates whether a reminder has been set for an item in the Exchange store.
	ReminderIsSet *ReminderIsSet `xml:"ReminderIsSet"`
	// The ReminderMinutesBeforeStart element represents the number of minutes before an event occurs when a reminder is displayed.
	ReminderMinutesBeforeStart *ReminderMinutesBeforeStart `xml:"ReminderMinutesBeforeStart"`
	// The ReplyTo element identifies an array of addresses to which replies should be sent.
	ReplyTo *ReplyTo `xml:"ReplyTo"`
	// The RequiredAttendees element represents attendees that are required to attend a meeting.
	RequiredAttendees *RequiredAttendees `xml:"RequiredAttendees"`
	// The Resources element represents a scheduled resource for a meeting.
	Resources *Resources `xml:"Resources"`
	// The ResponseObjects element contains a collection of all the response objects that are associated with an item in the Exchange store.
	ResponseObjects *ResponseObjects `xml:"ResponseObjects"`
	// The ResponseType element represents the type of recipient response that is received for a meeting.
	ResponseType *ResponseType `xml:"ResponseType"`
	// The Sender element identifies the sender of an item.
	Sender *Sender `xml:"Sender"`
	// The Sensitivity element indicates the sensitivity level of an item.
	Sensitivity *Sensitivity `xml:"Sensitivity"`
	// The Size element represents the size in bytes of an item or all the items in a conversation in the current folder. This property is read-only.
	Size *Size `xml:"Size"`
	// The Start element represents the start of a duration.
	Start *Start `xml:"Start"`
	// The StartTimeZone element defines the time zone for the start time of a CalendarItem or MeetingRequest.
	StartTimeZone *StartTimeZone `xml:"StartTimeZone"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The TimeZone element provides a text description of a time zone.
	TimeZone *TimeZoneItem `xml:"TimeZone"`
	// The ToRecipients element contains an array of recipients of an item. These are the primary recipients of an item.
	ToRecipients *ToRecipients `xml:"ToRecipients"`
	// The UID element uniquely identifies a calendar item.
	UID *UID `xml:"UID"`
	// The UniqueBody element represents an HTML fragment or plain text which represents the unique body of this conversation.
	UniqueBody *UniqueBody `xml:"UniqueBody"`
	// The WebClientEditFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to edit an item in Outlook Web App.
	WebClientEditFormQueryString *WebClientEditFormQueryString `xml:"WebClientEditFormQueryString"`
	// The WebClientReadFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to read an item in Outlook Web App.
	WebClientReadFormQueryString *WebClientReadFormQueryString `xml:"WebClientReadFormQueryString"`
	// The When element provides information about when a calendar or meeting item occurs.
	When *When `xml:"When"`
}

func (M *MeetingRequest) SetForMarshal() {
	M.XMLName.Local = "t:MeetingRequest"
}

func (M *MeetingRequest) GetSchema() *Schema {
	return &SchemaTypes
}
