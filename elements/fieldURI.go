package elements

// The FieldURIattr element identifies frequently referenced properties by URI.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri
type FieldURI struct {
	FieldURI *FieldURIattrFieldURI `xml:"FieldURI,attr"`
}

type FieldURIattrFieldURI string

const (
	// Identifies the FolderId property.
	FieldURIattrFieldURIfolderFolderId FieldURIattrFieldURI = `folder:FolderId`
	// Identifies the ParentFolderId property.
	FieldURIattrFieldURIfolderParentFolderId FieldURIattrFieldURI = `folder:ParentFolderId`
	// Identifies the DisplayName property.
	FieldURIattrFieldURIfolderDisplayName FieldURIattrFieldURI = `folder:DisplayName`
	// Identifies the UnreadCount property.
	FieldURIattrFieldURIfolderUnreadCount FieldURIattrFieldURI = `folder:UnreadCount`
	// Identifies the TotalCount property.
	FieldURIattrFieldURIfolderTotalCount FieldURIattrFieldURI = `folder:TotalCount`
	// Identifies the ChildFolderCount property.
	FieldURIattrFieldURIfolderChildFolderCount FieldURIattrFieldURI = `folder:ChildFolderCount`
	// Identifies the FolderClass property.
	FieldURIattrFieldURIfolderFolderClass FieldURIattrFieldURI = `folder:FolderClass`
	// Identifies the SearchParameters property.
	FieldURIattrFieldURIfolderSearchParameters FieldURIattrFieldURI = `folder:SearchParameters`
	// Identifies the ManagedFolderInformation property.
	FieldURIattrFieldURIfolderManagedFolderInformation FieldURIattrFieldURI = `folder:ManagedFolderInformation`
	// Identifies the PermissionSet property.
	FieldURIattrFieldURIfolderPermissionSet FieldURIattrFieldURI = `folder:PermissionSet`
	// Identifies the EffectiveRights property.
	FieldURIattrFieldURIfolderEffectiveRights FieldURIattrFieldURI = `folder:EffectiveRights`
	// Identifies the SharingEffectiveRights property.
	FieldURIattrFieldURIfolderSharingEffectiveRights FieldURIattrFieldURI = `folder:SharingEffectiveRights`
	// Identifies the ItemId property.
	FieldURIattrFieldURIitemItemId FieldURIattrFieldURI = `item:ItemId`
	// Identifies the ParentFolderId property.
	FieldURIattrFieldURIitemParentFolderId FieldURIattrFieldURI = `item:ParentFolderId`
	// Identifies the ItemClass property.
	FieldURIattrFieldURIitemItemClass FieldURIattrFieldURI = `item:ItemClass`
	// Identifies the MimeContent property.
	FieldURIattrFieldURIitemMimeContent FieldURIattrFieldURI = `item:MimeContent`
	// Identifies the Attachments property.
	FieldURIattrFieldURIitemAttachments FieldURIattrFieldURI = `item:Attachments`
	// Identifies the Subject property.
	FieldURIattrFieldURIitemSubject FieldURIattrFieldURI = `item:Subject`
	// Identifies the DateTimeReceived property.
	FieldURIattrFieldURIitemDateTimeReceived FieldURIattrFieldURI = `item:DateTimeReceived`
	// Identifies the Size property.
	FieldURIattrFieldURIitemSize FieldURIattrFieldURI = `item:Size`
	// Identifies the Categories property.
	FieldURIattrFieldURIitemCategories FieldURIattrFieldURI = `item:Categories`
	// Identifies the HasAttachments property.
	FieldURIattrFieldURIitemHasAttachments FieldURIattrFieldURI = `item:HasAttachments`
	// Identifies the Importance property.
	FieldURIattrFieldURIitemImportance FieldURIattrFieldURI = `item:Importance`
	// Identifies the InReplyTo property.
	FieldURIattrFieldURIitemInReplyTo FieldURIattrFieldURI = `item:InReplyTo`
	// Identifies the InternetMessageHeaders property.
	FieldURIattrFieldURIitemInternetMessageHeaders FieldURIattrFieldURI = `item:InternetMessageHeaders`
	// Identifies the IsAssociated property.
	FieldURIattrFieldURIitemIsAssociated FieldURIattrFieldURI = `item:IsAssociated`
	// Identifies the IsDraft property.
	FieldURIattrFieldURIitemIsDraft FieldURIattrFieldURI = `item:IsDraft`
	// Identifies the IsFromMe property.
	FieldURIattrFieldURIitemIsFromMe FieldURIattrFieldURI = `item:IsFromMe`
	// Identifies the IsResend property.
	FieldURIattrFieldURIitemIsResend FieldURIattrFieldURI = `item:IsResend`
	// Identifies the IsSubmitted property.
	FieldURIattrFieldURIitemIsSubmitted FieldURIattrFieldURI = `item:IsSubmitted`
	// Identifies the IsUnmodified property.
	FieldURIattrFieldURIitemIsUnmodified FieldURIattrFieldURI = `item:IsUnmodified`
	// Identifies the DateTimeSent property.
	FieldURIattrFieldURIitemDateTimeSent FieldURIattrFieldURI = `item:DateTimeSent`
	// Identifies the DateTimeCreated property.
	FieldURIattrFieldURIitemDateTimeCreated FieldURIattrFieldURI = `item:DateTimeCreated`
	// Identifies the Body property.
	FieldURIattrFieldURIitemBody FieldURIattrFieldURI = `item:Body`
	// Identifies the ResponseObjects property.
	FieldURIattrFieldURIitemResponseObjects FieldURIattrFieldURI = `item:ResponseObjects`
	// Identifies the Sensitivity property.
	FieldURIattrFieldURIitemSensitivity FieldURIattrFieldURI = `item:Sensitivity`
	// Identifies the ReminderDueBy property.
	FieldURIattrFieldURIitemReminderDueBy FieldURIattrFieldURI = `item:ReminderDueBy`
	// Identifies the ReminderIsSet property.
	FieldURIattrFieldURIitemReminderIsSet FieldURIattrFieldURI = `item:ReminderIsSet`
	// Identifies the ReminderNextTime property.
	FieldURIattrFieldURIitemReminderNextTime FieldURIattrFieldURI = `item:ReminderNextTime`
	// Identifies the ReminderMinutesBeforeStart property.
	FieldURIattrFieldURIitemReminderMinutesBeforeStart FieldURIattrFieldURI = `item:ReminderMinutesBeforeStart`
	// Identifies the DisplayTo property.
	FieldURIattrFieldURIitemDisplayTo FieldURIattrFieldURI = `item:DisplayTo`
	// Identifies the DisplayCc property.
	FieldURIattrFieldURIitemDisplayCc FieldURIattrFieldURI = `item:DisplayCc`
	// Identifies the Culture property.
	FieldURIattrFieldURIitemCulture FieldURIattrFieldURI = `item:Culture`
	// Identifies the EffectiveRights property.
	FieldURIattrFieldURIitemEffectiveRights FieldURIattrFieldURI = `item:EffectiveRights`
	// Identifies the LastModifiedName property.
	FieldURIattrFieldURIitemLastModifiedName FieldURIattrFieldURI = `item:LastModifiedName`
	// Identifies the LastModifiedTime property.
	FieldURIattrFieldURIitemLastModifiedTime FieldURIattrFieldURI = `item:LastModifiedTime`
	// Identifies the ConversationId property.
	FieldURIattrFieldURIitemConversationId FieldURIattrFieldURI = `item:ConversationId`
	// Identifies the UniqueBody property.
	FieldURIattrFieldURIitemUniqueBody FieldURIattrFieldURI = `item:UniqueBody`
	// Identifies the Flag property.
	FieldURIattrFieldURIitemFlag FieldURIattrFieldURI = `item:Flag`
	// Identifies the StoreEntryId property.
	FieldURIattrFieldURIitemStoreEntryId FieldURIattrFieldURI = `item:StoreEntryId`
	// Identifies the InstanceKey property.
	FieldURIattrFieldURIitemInstanceKey FieldURIattrFieldURI = `item:InstanceKey`
	// Identifies the NormalizedBody property.
	FieldURIattrFieldURIitemNormalizedBody FieldURIattrFieldURI = `item:NormalizedBody`
	// Identifies the EntityExtractionResult property.
	FieldURIattrFieldURIitemEntityExtractionResult FieldURIattrFieldURI = `item:EntityExtractionResult`
	// Idnentifies the PolicyTag property.
	FieldURIattrFieldURIitemPolicyTag FieldURIattrFieldURI = `itemPolicyTag`
	// Identifies the ArchiveTag property.
	FieldURIattrFieldURIitemArchiveTag FieldURIattrFieldURI = `item:ArchiveTag`
	// Identifies the RetentionDate property.
	FieldURIattrFieldURIitemRetentionDate FieldURIattrFieldURI = `item:RetentionDate`
	// Identifies the Preview property.
	FieldURIattrFieldURIitemPreview FieldURIattrFieldURI = `item:Preview`
	// Identifies the NextPredictedAction property.
	FieldURIattrFieldURIitemNextPredictedAction FieldURIattrFieldURI = `item:NextPredictedAction`
	// Identifies the GroupingAction property.
	FieldURIattrFieldURIitemGroupingAction FieldURIattrFieldURI = `item:GroupingAction`
	// Identifies the PredictedActionReasons property
	FieldURIattrFieldURIitemPredictedActionReasons FieldURIattrFieldURI = `item:PredictedActionReasons`
	// Intended for internal use only.
	FieldURIattrFieldURIitemIsClutter FieldURIattrFieldURI = `item:IsClutter`
	// Identifies the RightsManagementLicenseData property.
	FieldURIattrFieldURIitemRightsManagementLicenseData FieldURIattrFieldURI = `item:RightsManagementLicenseData`
	// Identifies the BlockStatus property.
	FieldURIattrFieldURIitemBlockStatus FieldURIattrFieldURI = `item:BlockStatus`
	// Identifies the HasBlockedImages property.
	FieldURIattrFieldURIitemHasBlockedImages FieldURIattrFieldURI = `item:HasBlockedImages`
	// Identifies the WebClientReadFormQueryString property.
	FieldURIattrFieldURIitemWebClientReadFormQueryString FieldURIattrFieldURI = `item:WebClientReadFormQueryString`
	// Identifies the WebClientEditFormQueryString property.
	FieldURIattrFieldURIitemWebClientEditFormQueryString FieldURIattrFieldURI = `item:WebClientEditFormQueryString`
	// Identifies the TextBody property.
	FieldURIattrFieldURIitemTextBody FieldURIattrFieldURI = `item:TextBody`
	// Identifies the IconIndex property.
	FieldURIattrFieldURIitemIconIndex FieldURIattrFieldURI = `item:IconIndex`
	// Identifies the MimeContentUTF8 property.
	FieldURIattrFieldURIitemMimeContentUTF8 FieldURIattrFieldURI = `item:MimeContentUTF8`
	// Identifies the ConversationIndex property.
	FieldURIattrFieldURImessageConversationIndex FieldURIattrFieldURI = `message:ConversationIndex`
	// Identifies the ConversationTopic property.
	FieldURIattrFieldURImessageConversationTopic FieldURIattrFieldURI = `message:ConversationTopic`
	// Identifies the InternetMessageId property.
	FieldURIattrFieldURImessageInternetMessageId FieldURIattrFieldURI = `message:InternetMessageId`
	// Identifies the IsRead property.
	FieldURIattrFieldURImessageIsRead FieldURIattrFieldURI = `message:IsRead`
	// Identifies the IsResponseRequested property.
	FieldURIattrFieldURImessageIsResponseRequested FieldURIattrFieldURI = `message:IsResponseRequested`
	// Identifies the IsReadReceiptRequested property.
	FieldURIattrFieldURImessageIsReadReceiptRequested FieldURIattrFieldURI = `message:IsReadReceiptRequested`
	// Identifies the IsDeliveryReceiptRequested property.
	FieldURIattrFieldURImessageIsDeliveryReceiptRequested FieldURIattrFieldURI = `message:IsDeliveryReceiptRequested`
	// Identifies the ReceivedBy property.
	FieldURIattrFieldURImessageReceivedBy FieldURIattrFieldURI = `message:ReceivedBy`
	// Identifies the ReceivedRepresenting property.
	FieldURIattrFieldURImessageReceivedRepresenting FieldURIattrFieldURI = `message:ReceivedRepresenting`
	// Identifies the References property.
	FieldURIattrFieldURImessageReferences FieldURIattrFieldURI = `message:References`
	// Identifies the ReplyTo property.
	FieldURIattrFieldURImessageReplyTo FieldURIattrFieldURI = `message:ReplyTo`
	// Identifies the From property.
	FieldURIattrFieldURImessageFrom FieldURIattrFieldURI = `message:From`
	// Identifies the Sender property.
	FieldURIattrFieldURImessageSender FieldURIattrFieldURI = `message:Sender`
	// Identifies the ToRecipients property.
	FieldURIattrFieldURImessageToRecipients FieldURIattrFieldURI = `message:ToRecipients`
	// Identifies the CcRecipients property.
	FieldURIattrFieldURImessageCcRecipients FieldURIattrFieldURI = `message:CcRecipients`
	// Identifies the BccRecipients property.
	FieldURIattrFieldURImessageBccRecipients FieldURIattrFieldURI = `message:BccRecipients`
	// Identifies the ApprovalRequestData property.
	FieldURIattrFieldURImessageApprovalRequestData FieldURIattrFieldURI = `message:ApprovalRequestData`
	// Identifies the VotingInformation property.
	FieldURIattrFieldURImessageVotingInformation FieldURIattrFieldURI = `message:VotingInformation`
	// Identifies the ReminderMessageData property.
	FieldURIattrFieldURImessageReminderMessageData FieldURIattrFieldURI = `message:ReminderMessageData`
	// Identifies the AssociatedCalendarItemId property.
	FieldURIattrFieldURImeetingAssociatedCalendarItemId FieldURIattrFieldURI = `meeting:AssociatedCalendarItemId`
	// Identifies the IsDelegated property.
	FieldURIattrFieldURImeetingIsDelegated FieldURIattrFieldURI = `meeting:IsDelegated`
	// Identifies the IsOutOfDate property.
	FieldURIattrFieldURImeetingIsOutOfDate FieldURIattrFieldURI = `meeting:IsOutOfDate`
	// Identifies the HasBeenProcessed property.
	FieldURIattrFieldURImeetingHasBeenProcessed FieldURIattrFieldURI = `meeting:HasBeenProcessed`
	// Identifies the ResponseType property.
	FieldURIattrFieldURImeetingResponseType FieldURIattrFieldURI = `meeting:ResponseType`
	// Identifies the ProposedStart property.
	FieldURIattrFieldURImeetingProposedStart FieldURIattrFieldURI = `meeting:ProposedStart`
	// Identifies the ProposedEnd property.
	FieldURIattrFieldURImeetingProposedEnd FieldURIattrFieldURI = `meeting:ProposedEnd`
	// Identifies the MeetingRequestType property.
	FieldURIattrFieldURImeetingRequestMeetingRequestType FieldURIattrFieldURI = `meetingRequest:MeetingRequestType`
	// Identifies the IntendedFreeBusyStatus property.
	FieldURIattrFieldURImeetingRequestIntendedFreeBusyStatus FieldURIattrFieldURI = `meetingRequest:IntendedFreeBusyStatus`
	// Identifies the ChangeHighlights property.
	FieldURIattrFieldURImeetingRequestChangeHighlights FieldURIattrFieldURI = `meetingRequest:ChangeHighlights`
	// Identifies the Start property.
	FieldURIattrFieldURIcalendarStart FieldURIattrFieldURI = `calendar:Start`
	// Identifies the End property.
	FieldURIattrFieldURIcalendarEnd FieldURIattrFieldURI = `calendar:End`
	// Identifies the OriginalStart property.
	FieldURIattrFieldURIcalendarOriginalStart FieldURIattrFieldURI = `calendar:OriginalStart`
	// Identifies the StartWallClock property.
	FieldURIattrFieldURIcalendarStartWallClock FieldURIattrFieldURI = `calendar:StartWallClock`
	// Identifies the EndWallClock property.
	FieldURIattrFieldURIcalendarEndWallClock FieldURIattrFieldURI = `calendar:EndWallClock`
	// Identifies the StartTimeZoneId property.
	FieldURIattrFieldURIcalendarStartTimeZoneId FieldURIattrFieldURI = `calendar:StartTimeZoneId`
	// Identifies the EndTimeZoneId property.
	FieldURIattrFieldURIcalendarEndTimeZoneId FieldURIattrFieldURI = `calendar:EndTimeZoneId`
	// Identifies the IsAllDayEvent property.
	FieldURIattrFieldURIcalendarIsAllDayEvent FieldURIattrFieldURI = `calendar:IsAllDayEvent`
	// Identifies the LegacyFreeBusyStatus property.
	FieldURIattrFieldURIcalendarLegacyFreeBusyStatus FieldURIattrFieldURI = `calendar:LegacyFreeBusyStatus`
	// Identifies the Location property.
	FieldURIattrFieldURIcalendarLocation FieldURIattrFieldURI = `calendar:Location`
	// Identifies the When property.
	FieldURIattrFieldURIcalendarWhen FieldURIattrFieldURI = `calendar:When`
	// Identifies the IsMeeting property.
	FieldURIattrFieldURIcalendarIsMeeting FieldURIattrFieldURI = `calendar:IsMeeting`
	// Identifies the IsCancelled property.
	FieldURIattrFieldURIcalendarIsCancelled FieldURIattrFieldURI = `calendar:IsCancelled`
	// Identifies the IsRecurring property.
	FieldURIattrFieldURIcalendarIsRecurring FieldURIattrFieldURI = `calendar:IsRecurring`
	// Identifies the MeetingRequestWasSent property.
	FieldURIattrFieldURIcalendarMeetingRequestWasSent FieldURIattrFieldURI = `calendar:MeetingRequestWasSent`
	// Identifies the IsResponseRequested property.
	FieldURIattrFieldURIcalendarIsResponseRequested FieldURIattrFieldURI = `calendar:IsResponseRequested`
	// Identifies the CalendarItemType property.
	FieldURIattrFieldURIcalendarCalendarItemType FieldURIattrFieldURI = `calendar:CalendarItemType`
	// Identifies the MyResponseType property.
	FieldURIattrFieldURIcalendarMyResponseType FieldURIattrFieldURI = `calendar:MyResponseType`
	// Identifies the Organizer property.
	FieldURIattrFieldURIcalendarOrganizer FieldURIattrFieldURI = `calendar:Organizer`
	// Identifies the RequiredAttendees property.
	FieldURIattrFieldURIcalendarRequiredAttendees FieldURIattrFieldURI = `calendar:RequiredAttendees`
	// Identifies the OptionalAttendees property.
	FieldURIattrFieldURIcalendarOptionalAttendees FieldURIattrFieldURI = `calendar:OptionalAttendees`
	// Identifies the Resources property.
	FieldURIattrFieldURIcalendarResources FieldURIattrFieldURI = `calendar:Resources`
	// Identifies the ConflictingMeetingCount property.
	FieldURIattrFieldURIcalendarConflictingMeetingCount FieldURIattrFieldURI = `calendar:ConflictingMeetingCount`
	// Identifies the AdjacentMeetingCount property.
	FieldURIattrFieldURIcalendarAdjacentMeetingCount FieldURIattrFieldURI = `calendar:AdjacentMeetingCount`
	// Identifies the ConflictingMeetings property.
	FieldURIattrFieldURIcalendarConflictingMeetings FieldURIattrFieldURI = `calendar:ConflictingMeetings`
	// Identifies the AdjacentMeetings property.
	FieldURIattrFieldURIcalendarAdjacentMeetings FieldURIattrFieldURI = `calendar:AdjacentMeetings`
	// Identifies the Duration property.
	FieldURIattrFieldURIcalendarDuration FieldURIattrFieldURI = `calendar:Duration`
	// Identifies the TimeZone property.
	FieldURIattrFieldURIcalendarTimeZone FieldURIattrFieldURI = `calendar:TimeZone`
	// Identifies the AppointmentReplyTime property.
	FieldURIattrFieldURIcalendarAppointmentReplyTime FieldURIattrFieldURI = `calendar:AppointmentReplyTime`
	// Identifies the AppointmentSequenceNumber property.
	FieldURIattrFieldURIcalendarAppointmentSequenceNumber FieldURIattrFieldURI = `calendar:AppointmentSequenceNumber`
	// Identifies the AppointmentState property.
	FieldURIattrFieldURIcalendarAppointmentState FieldURIattrFieldURI = `calendar:AppointmentState`
	// Identifies the Recurrence property.
	FieldURIattrFieldURIcalendarRecurrence FieldURIattrFieldURI = `calendar:Recurrence`
	// Identifies the FirstOccurrence property.
	FieldURIattrFieldURIcalendarFirstOccurrence FieldURIattrFieldURI = `calendar:FirstOccurrence`
	// Identifies the LastOccurrence property.
	FieldURIattrFieldURIcalendarLastOccurrence FieldURIattrFieldURI = `calendar:LastOccurrence`
	// Identifies the ModifiedOccurrences property.
	FieldURIattrFieldURIcalendarModifiedOccurrences FieldURIattrFieldURI = `calendar:ModifiedOccurrences`
	// Identifies the DeletedOccurrences property.
	FieldURIattrFieldURIcalendarDeletedOccurrences FieldURIattrFieldURI = `calendar:DeletedOccurrences`
	// Identifies the MeetingTimeZone property.
	FieldURIattrFieldURIcalendarMeetingTimeZone FieldURIattrFieldURI = `calendar:MeetingTimeZone`
	// Identifies the ConferenceType property.
	FieldURIattrFieldURIcalendarConferenceType FieldURIattrFieldURI = `calendar:ConferenceType`
	// Identifies the AllowNewTimeProposal property.
	FieldURIattrFieldURIcalendarAllowNewTimeProposal FieldURIattrFieldURI = `calendar:AllowNewTimeProposal`
	// Identifies the IsOnlineMeeting property.
	FieldURIattrFieldURIcalendarIsOnlineMeeting FieldURIattrFieldURI = `calendar:IsOnlineMeeting`
	// Identifies the MeetingWorkspaceUrl property.
	FieldURIattrFieldURIcalendarMeetingWorkspaceUrl FieldURIattrFieldURI = `calendar:MeetingWorkspaceUrl`
	// Identifies the NetShowUrl property.
	FieldURIattrFieldURIcalendarNetShowUrl FieldURIattrFieldURI = `calendar:NetShowUrl`
	// Identifies the UID property.
	FieldURIattrFieldURIcalendarUID FieldURIattrFieldURI = `calendar:UID`
	// Identifies the RecurrenceId property.
	FieldURIattrFieldURIcalendarRecurrenceId FieldURIattrFieldURI = `calendar:RecurrenceId`
	// Identifies the DateTimeStamp property.
	FieldURIattrFieldURIcalendarDateTimeStamp FieldURIattrFieldURI = `calendar:DateTimeStamp`
	// Identifies the StartTimeZone property.
	FieldURIattrFieldURIcalendarStartTimeZone FieldURIattrFieldURI = `calendar:StartTimeZone`
	// Identifies the EndTimeZone property.
	FieldURIattrFieldURIcalendarEndTimeZone FieldURIattrFieldURI = `calendar:EndTimeZone`
	// Identifies the JoinOnlineMeetingUrl property.
	FieldURIattrFieldURIcalendarJoinOnlineMeetingUrl FieldURIattrFieldURI = `calendar:JoinOnlineMeetingUrl`
	// Identifies the OnlineMeetingSettings property.
	FieldURIattrFieldURIcalendarOnlineMeetingSettings FieldURIattrFieldURI = `calendar:OnlineMeetingSettings`
	// Identifies the IsOrganizer property.
	FieldURIattrFieldURIcalendarIsOrganizer FieldURIattrFieldURI = `calendar:IsOrganizer`
	// Identifies the ActualWork property.
	FieldURIattrFieldURItaskActualWork FieldURIattrFieldURI = `task:ActualWork`
	// Identifies the AssignedTime property.
	FieldURIattrFieldURItaskAssignedTime FieldURIattrFieldURI = `task:AssignedTime`
	// Identifies the BillingInformation property.
	FieldURIattrFieldURItaskBillingInformation FieldURIattrFieldURI = `task:BillingInformation`
	// Identifies the ChangeCount property.
	FieldURIattrFieldURItaskChangeCount FieldURIattrFieldURI = `task:ChangeCount`
	// Identifies the Companies property.
	FieldURIattrFieldURItaskCompanies FieldURIattrFieldURI = `task:Companies`
	// Identifies the CompleteDate property.
	FieldURIattrFieldURItaskCompleteDate FieldURIattrFieldURI = `task:CompleteDate`
	// Identifies the Contacts property.
	FieldURIattrFieldURItaskContacts FieldURIattrFieldURI = `task:Contacts`
	// Identifies the DelegationState property.
	FieldURIattrFieldURItaskDelegationState FieldURIattrFieldURI = `task:DelegationState`
	// Identifies the Delegator property.
	FieldURIattrFieldURItaskDelegator FieldURIattrFieldURI = `task:Delegator`
	// Identifies the DueDate property.
	FieldURIattrFieldURItaskDueDate FieldURIattrFieldURI = `task:DueDate`
	// Identifies the IsAssignmentEditable property.
	FieldURIattrFieldURItaskIsAssignmentEditable FieldURIattrFieldURI = `task:IsAssignmentEditable`
	// Identifies the IsComplete property.
	FieldURIattrFieldURItaskIsComplete FieldURIattrFieldURI = `task:IsComplete`
	// Identifies the IsRecurring property.
	FieldURIattrFieldURItaskIsRecurring FieldURIattrFieldURI = `task:IsRecurring`
	// Identifies the IsTeamTask property.
	FieldURIattrFieldURItaskIsTeamTask FieldURIattrFieldURI = `task:IsTeamTask`
	// Identifies the Mileage property.
	FieldURIattrFieldURItaskMileage FieldURIattrFieldURI = `task:Mileage`
	// Identifies the Owner property.
	FieldURIattrFieldURItaskOwner FieldURIattrFieldURI = `task:Owner`
	// Identifies the PercentComplete property.
	FieldURIattrFieldURItaskPercentComplete FieldURIattrFieldURI = `task:PercentComplete`
	// Identifies the Recurrence property.
	FieldURIattrFieldURItaskRecurrence FieldURIattrFieldURI = `task:Recurrence`
	// Identifies the StartDate property.
	FieldURIattrFieldURItaskStartDate FieldURIattrFieldURI = `task:StartDate`
	// Identifies the Status property.
	FieldURIattrFieldURItaskStatus FieldURIattrFieldURI = `task:Status`
	// Identifies the StatusDescription property.
	FieldURIattrFieldURItaskStatusDescription FieldURIattrFieldURI = `task:StatusDescription`
	// Identifies the TotalWork property.
	FieldURIattrFieldURItaskTotalWork FieldURIattrFieldURI = `task:TotalWork`
	// Identifies the Alias property. This property was introduced in Exchange Server 2010 Service Pack 2 (SP2).
	FieldURIattrFieldURIcontactsAlias FieldURIattrFieldURI = `contacts:Alias`
	// Identifies the AssistantName property.
	FieldURIattrFieldURIcontactsAssistantName FieldURIattrFieldURI = `contacts:AssistantName`
	// Identifies the Birthday property.
	FieldURIattrFieldURIcontactsBirthday FieldURIattrFieldURI = `contacts:Birthday`
	// Identifies the BusinessHomePage property.
	FieldURIattrFieldURIcontactsBusinessHomePage FieldURIattrFieldURI = `contacts:BusinessHomePage`
	// Identifies the Children property.
	FieldURIattrFieldURIcontactsChildren FieldURIattrFieldURI = `contacts:Children`
	// Identifies the Companies property.
	FieldURIattrFieldURIcontactsCompanies FieldURIattrFieldURI = `contacts:Companies`
	// Identifies the CompanyName property.
	FieldURIattrFieldURIcontactsCompanyName FieldURIattrFieldURI = `contacts:CompanyName`
	// Identifies the CompleteName property.
	FieldURIattrFieldURIcontactsCompleteName FieldURIattrFieldURI = `contacts:CompleteName`
	// Identifies the ContactSource property.
	FieldURIattrFieldURIcontactsContactSource FieldURIattrFieldURI = `contacts:ContactSource`
	// Identifies the Culture property.
	FieldURIattrFieldURIcontactsCulture FieldURIattrFieldURI = `contacts:Culture`
	// Identifies the Department property.
	FieldURIattrFieldURIcontactsDepartment FieldURIattrFieldURI = `contacts:Department`
	// Identifies the DisplayName property.
	FieldURIattrFieldURIcontactsDisplayName FieldURIattrFieldURI = `contacts:DisplayName`
	// Identifies the DirectoryId property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsDirectoryId FieldURIattrFieldURI = `contacts:DirectoryId`
	// Identifies the DirectReports property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsDirectReports FieldURIattrFieldURI = `contacts:DirectReports`
	// Identifies the EmailAddresses property.
	FieldURIattrFieldURIcontactsEmailAddresses FieldURIattrFieldURI = `contacts:EmailAddresses`
	// Identifies the FileAs property.
	FieldURIattrFieldURIcontactsFileAs FieldURIattrFieldURI = `contacts:FileAs`
	// Identifies the FileAsMapping property.
	FieldURIattrFieldURIcontactsFileAsMapping FieldURIattrFieldURI = `contacts:FileAsMapping`
	// Identifies the Generation property.
	FieldURIattrFieldURIcontactsGeneration FieldURIattrFieldURI = `contacts:Generation`
	// Identifies the GivenName property.
	FieldURIattrFieldURIcontactsGivenName FieldURIattrFieldURI = `contacts:GivenName`
	// Identifies the ImAddresses property.
	FieldURIattrFieldURIcontactsImAddresses FieldURIattrFieldURI = `contacts:ImAddresses`
	// Identifies the Initials property.
	FieldURIattrFieldURIcontactsInitials FieldURIattrFieldURI = `contacts:Initials`
	// Identifies the JobTitle property.
	FieldURIattrFieldURIcontactsJobTitle FieldURIattrFieldURI = `contacts:JobTitle`
	// Identifies the Manager property.
	FieldURIattrFieldURIcontactsManager FieldURIattrFieldURI = `contacts:Manager`
	// Identifies the ManagerMailbox property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsManagerMailbox FieldURIattrFieldURI = `contacts:ManagerMailbox`
	// Identifies the MiddleName property.
	FieldURIattrFieldURIcontactsMiddleName FieldURIattrFieldURI = `contacts:MiddleName`
	// Identifies the Mileage property.
	FieldURIattrFieldURIcontactsMileage FieldURIattrFieldURI = `contacts:Mileage`
	// Identifies the MSExchangeCertificate property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsMSExchangeCertificate FieldURIattrFieldURI = `contacts:MSExchangeCertificate`
	// Identifies the Nickname property.
	FieldURIattrFieldURIcontactsNickname FieldURIattrFieldURI = `contacts:Nickname`
	// Identifies the Notes property. This property was introduced with in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsNotes FieldURIattrFieldURI = `contacts:Notes`
	// Identifies the OfficeLocation property.
	FieldURIattrFieldURIcontactsOfficeLocation FieldURIattrFieldURI = `contacts:OfficeLocation`
	// Identifies the PhoneNumbers property.
	FieldURIattrFieldURIcontactsPhoneNumbers FieldURIattrFieldURI = `contacts:PhoneNumbers`
	// Identifies the PhoneticFullName property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsPhoneticFullName FieldURIattrFieldURI = `contacts:PhoneticFullName`
	// Identifies the PhoneticFirstName property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsPhoneticFirstName FieldURIattrFieldURI = `contacts:PhoneticFirstName`
	// Identifies the PhoneticLastName property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsPhoneticLastName FieldURIattrFieldURI = `contacts:PhoneticLastName`
	// Identifies the Photo property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsPhoto FieldURIattrFieldURI = `contacts:Photo`
	// Identifies the PhysicalAddresses property.
	FieldURIattrFieldURIcontactsPhysicalAddresses FieldURIattrFieldURI = `contacts:PhysicalAddresses`
	// Identifies the PostalAddressIndex property.
	FieldURIattrFieldURIcontactsPostalAddressIndex FieldURIattrFieldURI = `contacts:PostalAddressIndex`
	// Identifies the Profession property.
	FieldURIattrFieldURIcontactsProfession FieldURIattrFieldURI = `contacts:Profession`
	// Identifies the SpouseName property.
	FieldURIattrFieldURIcontactsSpouseName FieldURIattrFieldURI = `contacts:SpouseName`
	// Identifies the Surname property.
	FieldURIattrFieldURIcontactsSurname FieldURIattrFieldURI = `contacts:Surname`
	// Identifies the WeddingAnniversary property.
	FieldURIattrFieldURIcontactsWeddingAnniversary FieldURIattrFieldURI = `contacts:WeddingAnniversary`
	// Identifies the UserSMIMECertificate property. This property was introduced in Exchange 2010 SP2.
	FieldURIattrFieldURIcontactsUserSMIMECertificate FieldURIattrFieldURI = `contacts:UserSMIMECertificate`
	// Identifies the HasPicture property.
	FieldURIattrFieldURIcontactsHasPicture FieldURIattrFieldURI = `contacts:HasPicture`
	// Identifies the Members property.
	FieldURIattrFieldURIdistributionlistMembers FieldURIattrFieldURI = `distributionlist:Members`
	// Identifies the PostedTime property.
	FieldURIattrFieldURIpostitemPostedTime FieldURIattrFieldURI = `postitem:PostedTime`
	// Identifies the ConversationId property.
	FieldURIattrFieldURIconversationConversationId FieldURIattrFieldURI = `conversation:ConversationId`
	// Identifies the ConversationTopic property.
	FieldURIattrFieldURIconversationConversationTopic FieldURIattrFieldURI = `conversation:ConversationTopic`
	// Identifies the UniqueRecipients property.
	FieldURIattrFieldURIconversationUniqueRecipients FieldURIattrFieldURI = `conversation:UniqueRecipients`
	// Identifies the GlobalUniqueRecipients property.
	FieldURIattrFieldURIconversationGlobalUniqueRecipients FieldURIattrFieldURI = `conversation:GlobalUniqueRecipients`
	// Identifies the UniqueUnreadSenders property.
	FieldURIattrFieldURIconversationUniqueUnreadSenders FieldURIattrFieldURI = `conversation:UniqueUnreadSenders`
	// Identifies the GlobalUniqueUnreadSenders property.
	FieldURIattrFieldURIconversationGlobalUniqueUnreadSenders FieldURIattrFieldURI = `conversation:GlobalUniqueUnreadSenders`
	// Identifies the UniqueSenders property.
	FieldURIattrFieldURIconversationUniqueSenders FieldURIattrFieldURI = `conversation:UniqueSenders`
	// Identifies the GlobalUniqueSenders property.
	FieldURIattrFieldURIconversationGlobalUniqueSenders FieldURIattrFieldURI = `conversation:GlobalUniqueSenders`
	// Identifies the LastDeliveryTime property.
	FieldURIattrFieldURIconversationLastDeliveryTime FieldURIattrFieldURI = `conversation:LastDeliveryTime`
	// Identifies the GlobalLastDeliveryTime property.
	FieldURIattrFieldURIconversationGlobalLastDeliveryTime FieldURIattrFieldURI = `conversation:GlobalLastDeliveryTime`
	// Identifies the Categories property.
	FieldURIattrFieldURIconversationCategories FieldURIattrFieldURI = `conversation:Categories`
	// Identifies the GlobalCategories property.
	FieldURIattrFieldURIconversationGlobalCategories FieldURIattrFieldURI = `conversation:GlobalCategories`
	// Identifies the FlagStatus property.
	FieldURIattrFieldURIconversationFlagStatus FieldURIattrFieldURI = `conversation:FlagStatus`
	// Identifies the GlobalFlagStatus property.
	FieldURIattrFieldURIconversationGlobalFlagStatus FieldURIattrFieldURI = `conversation:GlobalFlagStatus`
	// Identifies the HasAttachments property.
	FieldURIattrFieldURIconversationHasAttachments FieldURIattrFieldURI = `conversation:HasAttachments`
	// Identifies the GlobalHasAttachments property.
	FieldURIattrFieldURIconversationGlobalHasAttachments FieldURIattrFieldURI = `conversation:GlobalHasAttachments`
	// Identifies the HasIrm property.
	FieldURIattrFieldURIconversationHasIrm FieldURIattrFieldURI = `conversation:HasIrm`
	// Identifies the GlobalHasIrm property.
	FieldURIattrFieldURIconversationGlobalHasIrm FieldURIattrFieldURI = `conversation:GlobalHasIrm`
	// Identifies the MessageCount property.
	FieldURIattrFieldURIconversationMessageCount FieldURIattrFieldURI = `conversation:MessageCount`
	// Identifies the GlobalMessageCount property.
	FieldURIattrFieldURIconversationGlobalMessageCount FieldURIattrFieldURI = `conversation:GlobalMessageCount`
	// Identifies the UnreadCount property.
	FieldURIattrFieldURIconversationUnreadCount FieldURIattrFieldURI = `conversation:UnreadCount`
	// Identifies the GlobalUnreadCount property.
	FieldURIattrFieldURIconversationGlobalUnreadCount FieldURIattrFieldURI = `conversation:GlobalUnreadCount`
	// Identifies the Size property.
	FieldURIattrFieldURIconversationSize FieldURIattrFieldURI = `conversation:Size`
	// Identifies the GlobalSize property.
	FieldURIattrFieldURIconversationGlobalSize FieldURIattrFieldURI = `conversation:GlobalSize`
	// Identifies the ItemClasses property.
	FieldURIattrFieldURIconversationItemClasses FieldURIattrFieldURI = `conversation:ItemClasses`
	// Identifies the GlobalItemClasses property.
	FieldURIattrFieldURIconversationGlobalItemClasses FieldURIattrFieldURI = `conversation:GlobalItemClasses`
	// Identifies the Importance property.
	FieldURIattrFieldURIconversationImportance FieldURIattrFieldURI = `conversation:Importance`
	// Identifies the GlobalImportance property.
	FieldURIattrFieldURIconversationGlobalImportance FieldURIattrFieldURI = `conversation:GlobalImportance`
	// Identifies the ItemIds property.
	FieldURIattrFieldURIconversationItemIds FieldURIattrFieldURI = `conversation:ItemIds`
	// Identifies the GlobalItemIds property.
	FieldURIattrFieldURIconversationGlobalItemIds FieldURIattrFieldURI = `conversation:GlobalItemIds`
	// Identifies the LastModifiedTime property.
	FieldURIattrFieldURIconversationLastModifiedTime FieldURIattrFieldURI = `conversation:LastModifiedTime`
	// Identifies the InstanceKey property.
	FieldURIattrFieldURIconversationInstanceKey FieldURIattrFieldURI = `conversation:InstanceKey`
	// Identifies the Preview property.
	FieldURIattrFieldURIconversationPreview FieldURIattrFieldURI = `conversation:Preview`
	// Identifies the GlobalParentFolderId property.
	FieldURIattrFieldURIconversationGlobalParentFolderId FieldURIattrFieldURI = `conversation:GlobalParentFolderId`
	// Identifies the NextPredictedAction property.
	FieldURIattrFieldURIconversationNextPredictedAction FieldURIattrFieldURI = `conversation:NextPredictedAction`
	// Identifies the GroupingAction property.
	FieldURIattrFieldURIconversationGroupingAction FieldURIattrFieldURI = `conversation:GroupingAction`
	// Identifies the IconIndex property.
	FieldURIattrFieldURIconversationIconIndex FieldURIattrFieldURI = `conversation:IconIndex`
	// Identifies the GlobalIconIndex property.
	FieldURIattrFieldURIconversationGlobalIconIndex FieldURIattrFieldURI = `conversation:GlobalIconIndex`
	// Identifies the DraftItemIds property.
	FieldURIattrFieldURIconversationDraftItemIds FieldURIattrFieldURI = `conversation:DraftItemIds`
	// Intended for internal use only.
	FieldURIattrFieldURIconversationHasClutter FieldURIattrFieldURI = `conversation:HasClutter`
	// Identifies the PersonaId property.
	FieldURIattrFieldURIpersonaPersonaId FieldURIattrFieldURI = `persona:PersonaId`
	// Identifies the PersonaType property.
	FieldURIattrFieldURIpersonaPersonaType FieldURIattrFieldURI = `persona:PersonaType`
	// Identifies the GivenName property.
	FieldURIattrFieldURIpersonaGivenName FieldURIattrFieldURI = `persona:GivenName`
	// Identifies the CompanyName property.
	FieldURIattrFieldURIpersonaCompanyName FieldURIattrFieldURI = `persona:CompanyName`
	// Identifies the Surname property.
	FieldURIattrFieldURIpersonaSurname FieldURIattrFieldURI = `persona:Surname`
	// Identifies the DisplayName property.
	FieldURIattrFieldURIpersonaDisplayName FieldURIattrFieldURI = `persona:DisplayName`
	// Identifies the EmailAddress property.
	FieldURIattrFieldURIpersonaEmailAddress FieldURIattrFieldURI = `persona:EmailAddress`
	// Identifies the FileAs property.
	FieldURIattrFieldURIpersonaFileAs FieldURIattrFieldURI = `persona:FileAs`
	// Identifies the HomeCity property.
	FieldURIattrFieldURIpersonaHomeCity FieldURIattrFieldURI = `persona:HomeCity`
	// Identifies the CreationTime property.
	FieldURIattrFieldURIpersonaCreationTime FieldURIattrFieldURI = `persona:CreationTime`
	// Identifies the RelevanceScore property.
	FieldURIattrFieldURIpersonaRelevanceScore FieldURIattrFieldURI = `persona:RelevanceScore`
	// Identifies the WorkCity property.
	FieldURIattrFieldURIpersonaWorkCity FieldURIattrFieldURI = `persona:WorkCity`
	// Identifies the PersonaObjectStatus property.
	FieldURIattrFieldURIpersonaPersonaObjectStatus FieldURIattrFieldURI = `persona:PersonaObjectStatus`
	// Identifies the FileAsId property.
	FieldURIattrFieldURIpersonaFileAsId FieldURIattrFieldURI = `persona:FileAsId`
	// Identifies the DisplayNamePrefix property.
	FieldURIattrFieldURIpersonaDisplayNamePrefix FieldURIattrFieldURI = `persona:DisplayNamePrefix`
	// Identifies the YomiCompanyName property.
	FieldURIattrFieldURIpersonaYomiCompanyName FieldURIattrFieldURI = `persona:YomiCompanyName`
	// Identifies the YomiFirstName property.
	FieldURIattrFieldURIpersonaYomiFirstName FieldURIattrFieldURI = `persona:YomiFirstName`
	// Identifies the YomiLastName property.
	FieldURIattrFieldURIpersonaYomiLastName FieldURIattrFieldURI = `persona:YomiLastName`
	// Identifies the Title property.
	FieldURIattrFieldURIpersonaTitle FieldURIattrFieldURI = `persona:Title`
	// Identifies the EmailAddresses property.
	FieldURIattrFieldURIpersonaEmailAddresses FieldURIattrFieldURI = `persona:EmailAddresses`
	// Identifies the PhoneNumber property.
	FieldURIattrFieldURIpersonaPhoneNumber FieldURIattrFieldURI = `persona:PhoneNumber`
	// Identifies the ImAddress property.
	FieldURIattrFieldURIpersonaImAddress FieldURIattrFieldURI = `persona:ImAddress`
	// Identifies the ImAddresses property.
	FieldURIattrFieldURIpersonaImAddresses FieldURIattrFieldURI = `persona:ImAddresses`
	// Identifies the ImAddresses2 property.
	FieldURIattrFieldURIpersonaImAddresses2 FieldURIattrFieldURI = `persona:ImAddresses2`
	// Identifies the ImAddresses3 property.
	FieldURIattrFieldURIpersonaImAddresses3 FieldURIattrFieldURI = `persona:ImAddresses3`
	// Identifies the FolderIds property.
	FieldURIattrFieldURIpersonaFolderIds FieldURIattrFieldURI = `persona:FolderIds`
	// Identifies the Attributions property.
	FieldURIattrFieldURIpersonaAttributions FieldURIattrFieldURI = `persona:Attributions`
	// Identifies the DisplayNames property.
	FieldURIattrFieldURIpersonaDisplayNames FieldURIattrFieldURI = `persona:DisplayNames`
	// Identifies the Initials property.
	FieldURIattrFieldURIpersonaInitials FieldURIattrFieldURI = `persona:Initials`
	// Identifies the FileAses property.
	FieldURIattrFieldURIpersonaFileAses FieldURIattrFieldURI = `persona:FileAses`
	// Identifies the FileAsIds property.
	FieldURIattrFieldURIpersonaFileAsIds FieldURIattrFieldURI = `persona:FileAsIds`
	// Identifies the DisplayNamePrefixes property.
	FieldURIattrFieldURIpersonaDisplayNamePrefixes FieldURIattrFieldURI = `persona:DisplayNamePrefixes`
	// Identifies the GivenNames property.
	FieldURIattrFieldURIpersonaGivenNames FieldURIattrFieldURI = `persona:GivenNames`
	// Identifies the MiddleNames property.
	FieldURIattrFieldURIpersonaMiddleNames FieldURIattrFieldURI = `persona:MiddleNames`
	// Identifies the Surnames property.
	FieldURIattrFieldURIpersonaSurnames FieldURIattrFieldURI = `persona:Surnames`
	// Identifies the Generations property.
	FieldURIattrFieldURIpersonaGenerations FieldURIattrFieldURI = `persona:Generations`
	// Identifies the Nicknames property.
	FieldURIattrFieldURIpersonaNicknames FieldURIattrFieldURI = `persona:Nicknames`
	// Identifies the YomiCompanyNames property.
	FieldURIattrFieldURIpersonaYomiCompanyNames FieldURIattrFieldURI = `persona:YomiCompanyNames`
	// Identifies the YomiFirstNames property.
	FieldURIattrFieldURIpersonaYomiFirstNames FieldURIattrFieldURI = `persona:YomiFirstNames`
	// Identifies the YomiLastNames property.
	FieldURIattrFieldURIpersonaYomiLastNames FieldURIattrFieldURI = `persona:YomiLastNames`
	// Identifies the BusinessPhoneNumbers property.
	FieldURIattrFieldURIpersonaBusinessPhoneNumbers FieldURIattrFieldURI = `persona:BusinessPhoneNumbers`
	// Identifies the BusinessPhoneNumbers2 property.
	FieldURIattrFieldURIpersonaBusinessPhoneNumbers2 FieldURIattrFieldURI = `persona:BusinessPhoneNumbers2`
	// Identifies the HomePhones property.
	FieldURIattrFieldURIpersonaHomePhones FieldURIattrFieldURI = `persona:HomePhones`
	// Identifies the HomePhones2 property.
	FieldURIattrFieldURIpersonaHomePhones2 FieldURIattrFieldURI = `persona:HomePhones2`
	// Identifies the MobilePhones property.
	FieldURIattrFieldURIpersonaMobilePhones FieldURIattrFieldURI = `persona:MobilePhones`
	// Identifies the MobilePhones2 property.
	FieldURIattrFieldURIpersonaMobilePhones2 FieldURIattrFieldURI = `persona:MobilePhones2`
	// Identifies the AssistantPhoneNumbers property.
	FieldURIattrFieldURIpersonaAssistantPhoneNumbers FieldURIattrFieldURI = `persona:AssistantPhoneNumbers`
	// Identifies the CallbackPhones property.
	FieldURIattrFieldURIpersonaCallbackPhones FieldURIattrFieldURI = `persona:CallbackPhones`
	// Identifies the CarPhones property.
	FieldURIattrFieldURIpersonaCarPhones FieldURIattrFieldURI = `persona:CarPhones`
	// Identifies the HomeFaxes property.
	FieldURIattrFieldURIpersonaHomeFaxes FieldURIattrFieldURI = `persona:HomeFaxes`
	// Identifies the OrganizationMainPhones property.
	FieldURIattrFieldURIpersonaOrganizationMainPhones FieldURIattrFieldURI = `persona:OrganizationMainPhones`
	// Identifies the OtherFaxes property.
	FieldURIattrFieldURIpersonaOtherFaxes FieldURIattrFieldURI = `persona:OtherFaxes`
	// Identifies the OtherTelephones property.
	FieldURIattrFieldURIpersonaOtherTelephones FieldURIattrFieldURI = `persona:OtherTelephones`
	// Identifies the OtherPhones2 property.
	FieldURIattrFieldURIpersonaOtherPhones2 FieldURIattrFieldURI = `persona:OtherPhones2`
	// Identifies the Pagers property.
	FieldURIattrFieldURIpersonaPagers FieldURIattrFieldURI = `persona:Pagers`
	// Identifies the RadioPhones property.
	FieldURIattrFieldURIpersonaRadioPhones FieldURIattrFieldURI = `persona:RadioPhones`
	// Identifies the TelexNumbers property.
	FieldURIattrFieldURIpersonaTelexNumbers FieldURIattrFieldURI = `persona:TelexNumbers`
	// Identifies the WorkFaxes property.
	FieldURIattrFieldURIpersonaWorkFaxes FieldURIattrFieldURI = `persona:WorkFaxes`
	// Identifies the Emails1 property.
	FieldURIattrFieldURIpersonaEmails1 FieldURIattrFieldURI = `persona:Emails1`
	// Identifies the Emails2 property.
	FieldURIattrFieldURIpersonaEmails2 FieldURIattrFieldURI = `persona:Emails2`
	// Identifies the Emails3 property.
	FieldURIattrFieldURIpersonaEmails3 FieldURIattrFieldURI = `persona:Emails3`
	// Identifies the BusinessHomePages property.
	FieldURIattrFieldURIpersonaBusinessHomePages FieldURIattrFieldURI = `persona:BusinessHomePages`
	// Identifies the School property.
	FieldURIattrFieldURIpersonaSchool FieldURIattrFieldURI = `persona:School`
	// Identifies the PersonalHomePages property.
	FieldURIattrFieldURIpersonaPersonalHomePages FieldURIattrFieldURI = `persona:PersonalHomePages`
	// Identifies the OfficeLocations property.
	FieldURIattrFieldURIpersonaOfficeLocations FieldURIattrFieldURI = `persona:OfficeLocations`
	// Identifies the BusinessAddresses property.
	FieldURIattrFieldURIpersonaBusinessAddresses FieldURIattrFieldURI = `persona:BusinessAddresses`
	// Identifies the HomeAddresses property.
	FieldURIattrFieldURIpersonaHomeAddresses FieldURIattrFieldURI = `persona:HomeAddresses`
	// Identifies the OtherAddresses property.
	FieldURIattrFieldURIpersonaOtherAddresses FieldURIattrFieldURI = `persona:OtherAddresses`
	// Identifies the Titles property.
	FieldURIattrFieldURIpersonaTitles FieldURIattrFieldURI = `persona:Titles`
	// Identifies the Departments property.
	FieldURIattrFieldURIpersonaDepartments FieldURIattrFieldURI = `persona:Departments`
	// Identifies the CompanyNames property.
	FieldURIattrFieldURIpersonaCompanyNames FieldURIattrFieldURI = `persona:CompanyNames`
	// Identifies the Managers property.
	FieldURIattrFieldURIpersonaManagers FieldURIattrFieldURI = `persona:Managers`
	// Identifies the AssistantNames property.
	FieldURIattrFieldURIpersonaAssistantNames FieldURIattrFieldURI = `persona:AssistantNames`
	// Identifies the Professions property.
	FieldURIattrFieldURIpersonaProfessions FieldURIattrFieldURI = `persona:Professions`
	// Identifies the SpouseNames property.
	FieldURIattrFieldURIpersonaSpouseNames FieldURIattrFieldURI = `persona:SpouseNames`
	// Identifies the Hobbies property.
	FieldURIattrFieldURIpersonaHobbies FieldURIattrFieldURI = `persona:Hobbies`
	// Identifies the WeddingAnniversaries property.
	FieldURIattrFieldURIpersonaWeddingAnniversaries FieldURIattrFieldURI = `persona:WeddingAnniversaries`
	// Identifies the Birthdays property.
	FieldURIattrFieldURIpersonaBirthdays FieldURIattrFieldURI = `persona:Birthdays`
	// Identifies the Children property.
	FieldURIattrFieldURIpersonaChildren FieldURIattrFieldURI = `persona:Children`
	// Identifies the Locations property.
	FieldURIattrFieldURIpersonaLocations FieldURIattrFieldURI = `persona:Locations`
	// Identifies the ExtendedProperties property.
	FieldURIattrFieldURIpersonaExtendedProperties FieldURIattrFieldURI = `persona:ExtendedProperties`
	// Identifies the PostalAddress property.
	FieldURIattrFieldURIpersonaPostalAddress FieldURIattrFieldURI = `persona:PostalAddress`
	// Identifies the Bodies property.
	FieldURIattrFieldURIpersonaBodies FieldURIattrFieldURI = `persona:Bodies`
)
