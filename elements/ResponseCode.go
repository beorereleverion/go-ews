package elements

// The ResponseCode element provides status information about the request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/responsecode
import "encoding/xml"

type ResponseCode struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// This error occurs when the operation failed because of communication problems with Active Directory Domain Services (AD DS).
	ResponseCodeErrorADOperation string = `ErrorADOperation`
	// This error is returned when a ResolveNames operation request specifies a name that is not valid.
	ResponseCodeErrorADSessionFilter string = `ErrorADSessionFilter`
	// This error occurs when AD DS is unavailable. Try your request again later.
	ResponseCodeErrorADUnavailable string = `ErrorADUnavailable`
	// This error occurs when the calling account does not have the rights to perform the requested action.
	ResponseCodeErrorAccessDenied string = `ErrorAccessDenied`
	// This error is for internal use only. This error is not returned.
	ResponseCodeErrorAccessModeSpecified string = `ErrorAccessModeSpecified`
	// This error occurs when the account in question has been disabled.
	ResponseCodeErrorAccountDisabled string = `ErrorAccountDisabled`
	// This error occurs when a list with added delegates cannot be saved.
	ResponseCodeErrorAddDelegatesFailed string = `ErrorAddDelegatesFailed`
	// This error occurs when the address space record, or Domain Name System (DNS) domain name, for cross-forest availability could not be found in the Active Directory database.
	ResponseCodeErrorAddressSpaceNotFound string = `ErrorAddressSpaceNotFound`
	// This error indicates that the AffectedTaskOccurrences attribute was not specified. When the DeleteItem element is used to delete at least one item that is a task, and regardless of whether that task is recurring or not, the AffectedTaskOccurrences attribute has to be specified so that DeleteItem can determine whether to delete the current occurrence or the entire series.
	ResponseCodeErrorAffectedTaskOccurrencesRequired string = `ErrorAffectedTaskOccurrencesRequired`
	// This error MUST be returned if an action cannot be applied to one or more items in the conversation.
	ResponseCodeErrorApplyConversationActionFailed string = `ErrorApplyConversationActionFailed`
	// Indicates an error in archive folder path creation.
	ResponseCodeErrorArchiveFolderPathCreation string = `ErrorArchiveFolderPathCreation`
	// Indicates that the archive mailbox was not enabled.
	ResponseCodeErrorArchiveMailboxNotEnabled string = `ErrorArchiveMailboxNotEnabled`
	// This error is returned when an archive mailbox search is unsuccessful. This error was introduced in Exchange 2013.
	ResponseCodeErrorArchiveMailboxSearchFailed string = `ErrorArchiveMailboxSearchFailed`
	// This error is returned when the URL of an archive mailbox is not discoverable. This error was introduced in Exchange 2013.
	ResponseCodeErrorArchiveMailboxServiceDiscoveryFailed string = `ErrorArchiveMailboxServiceDiscoveryFailed`
	// Specifies that an attempt was made to create an item with more than 10 nested attachments. This value was introduced in Exchange Server 2010 Service Pack 2 (SP2).
	ResponseCodeErrorAttachmentNestLevelLimitExceeded string = `ErrorAttachmentNestLevelLimitExceeded`
	// The CreateAttachment element returns this error if an attempt is made to create an attachment with size exceeding Int32.MaxValue, in bytes. The GetAttachment element returns this error if an attempt to retrieve an existing attachment with size exceeding Int32.MaxValue, in bytes.
	ResponseCodeErrorAttachmentSizeLimitExceeded string = `ErrorAttachmentSizeLimitExceeded`
	// This error indicates that Exchange Web Services tried to determine the location of a cross-forest computer that is running Exchange 2010 that has the Client Access server role installed by using the Autodiscover service, but the call to the Autodiscover service failed.
	ResponseCodeErrorAutoDiscoverFailed string = `ErrorAutoDiscoverFailed`
	// This error indicates that the availability configuration information for the local forest is missing from AD DS.
	ResponseCodeErrorAvailabilityConfigNotFound string = `ErrorAvailabilityConfigNotFound`
	// This error indicates that an exception occurred while processing an item and that exception is likely to occur for the items that follow. Requests may include multiple items; for example, a GetItem operation request might include multiple identifiers. In general, items are processed one at a time. If an exception occurs while processing an item and that exception is likely to occur for the items that follow, items that follow will not be processed. The following are examples of errors that will stop processing for items that follow:  - ErrorAccessDenied  - ErrorAccountDisabled  - ErrorADUnavailable  - ErrorADOperation  - ErrorConnectionFailed  - ErrorMailboxStoreUnavailable  - ErrorMailboxMoveInProgress  - ErrorPasswordChangeRequired  - ErrorPasswordExpired  - ErrorQuotaExceeded  - ErrorInsufficientResources
	ResponseCodeErrorBatchProcessingStopped string = `ErrorBatchProcessingStopped`
	// This error occurs when an attempt is made to move or copy an occurrence of a recurring calendar item.
	ResponseCodeErrorCalendarCannotMoveOrCopyOccurrence string = `ErrorCalendarCannotMoveOrCopyOccurrence`
	// This error occurs when an attempt is made to update a calendar item that is located in the Deleted Items folder and when meeting updates or cancellations are to be sent according to the value of the SendMeetingInvitationsOrCancellations attribute. The following are the possible values for this attribute:  - SendToAllAndSaveCopy  - SendToChangedAndSaveCopy  - SendOnlyToAll  - SendOnlyToChanged  However, such an update is allowed only when the value of this attribute is set to SendToNone.
	ResponseCodeErrorCalendarCannotUpdateDeletedItem string = `ErrorCalendarCannotUpdateDeletedItem`
	// This error occurs when the UpdateItem, GetItem, DeleteItem, MoveItem, CopyItem, or SendItem operation is called and the ID that was specified is not an occurrence ID of any recurring calendar item.
	ResponseCodeErrorCalendarCannotUseIdForOccurrenceId string = `ErrorCalendarCannotUseIdForOccurrenceId`
	// This error occurs when the UpdateItem, GetItem, DeleteItem, MoveItem, CopyItem, or SendItem operation is called and the ID that was specified is not an ID of any recurring master item.
	ResponseCodeErrorCalendarCannotUseIdForRecurringMasterId string = `ErrorCalendarCannotUseIdForRecurringMasterId`
	// This error occurs during a CreateItem or UpdateItem operation when a calendar item duration is longer than the maximum allowed, which is currently 5 years.
	ResponseCodeErrorCalendarDurationIsTooLong string = `ErrorCalendarDurationIsTooLong`
	// This error occurs when a calendar End time is set to the same time or after the Start time.
	ResponseCodeErrorCalendarEndDateIsEarlierThanStartDate string = `ErrorCalendarEndDateIsEarlierThanStartDate`
	// This error occurs when the specified folder for a FindItem operation with a CalendarView element is not of calendar folder type.
	ResponseCodeErrorCalendarFolderIsInvalidForCalendarView string = `ErrorCalendarFolderIsInvalidForCalendarView`
	// This response code is not used.
	ResponseCodeErrorCalendarInvalidAttributeValue string = `ErrorCalendarInvalidAttributeValue`
	// This error occurs during a CreateItem or UpdateItem operation when invalid values of Day, WeekendDay, and Weekday are used to define the time change pattern.
	ResponseCodeErrorCalendarInvalidDayForTimeChangePattern string = `ErrorCalendarInvalidDayForTimeChangePattern`
	// This error occurs during a CreateItem or UpdateItem operation when invalid values of Day, WeekDay, and WeekendDay are used to specify the weekly recurrence.
	ResponseCodeErrorCalendarInvalidDayForWeeklyRecurrence string = `ErrorCalendarInvalidDayForWeeklyRecurrence`
	// This error occurs when the state of a calendar item recurrence binary large object (BLOB) in the Exchange store is invalid.
	ResponseCodeErrorCalendarInvalidPropertyState string = `ErrorCalendarInvalidPropertyState`
	// This response code is not used.
	ResponseCodeErrorCalendarInvalidPropertyValue string = `ErrorCalendarInvalidPropertyValue`
	// This error occurs when the specified recurrence cannot be created.
	ResponseCodeErrorCalendarInvalidRecurrence string = `ErrorCalendarInvalidRecurrence`
	// This error occurs when an invalid time zone is encountered.
	ResponseCodeErrorCalendarInvalidTimeZone string = `ErrorCalendarInvalidTimeZone`
	// This error Indicates that a calendar item has been canceled.
	ResponseCodeErrorCalendarIsCancelledForAccept string = `ErrorCalendarIsCancelledForAccept`
	// This error indicates that a calendar item has been canceled.
	ResponseCodeErrorCalendarIsCancelledForDecline string = `ErrorCalendarIsCancelledForDecline`
	// This error indicates that a calendar item has been canceled.
	ResponseCodeErrorCalendarIsCancelledForRemove string = `ErrorCalendarIsCancelledForRemove`
	// This error indicates that a calendar item has been canceled.
	ResponseCodeErrorCalendarIsCancelledForTentative string = `ErrorCalendarIsCancelledForTentative`
	// This error indicates that the AcceptItem element is invalid for a calendar item or meeting request in a delegated scenario.
	ResponseCodeErrorCalendarIsDelegatedForAccept string = `ErrorCalendarIsDelegatedForAccept`
	// This error indicates that the DeclineItem element is invalid for a calendar item or meeting request in a delegated scenario.
	ResponseCodeErrorCalendarIsDelegatedForDecline string = `ErrorCalendarIsDelegatedForDecline`
	// This error indicates that the RemoveItem element is invalid for a meeting cancellation in a delegated scenario.
	ResponseCodeErrorCalendarIsDelegatedForRemove string = `ErrorCalendarIsDelegatedForRemove`
	// This error indicates that the TentativelyAcceptItem element is invalid for a calendar item or meeting request in a delegated scenario.
	ResponseCodeErrorCalendarIsDelegatedForTentative string = `ErrorCalendarIsDelegatedForTentative`
	// This error is intended for internal use only.
	ResponseCodeErrorCalendarIsGroupMailboxForAccept string = `ErrorCalendarIsGroupMailboxForAccept`
	// This error is intended for internal use only.
	ResponseCodeErrorCalendarIsGroupMailboxForDecline string = `ErrorCalendarIsGroupMailboxForDecline`
	// This error is intended for internal use only.
	ResponseCodeErrorCalendarIsGroupMailboxForSuppressReadReceipt string = `ErrorCalendarIsGroupMailboxForSuppressReadReceipt`
	// This error is intended for internal use only.
	ResponseCodeErrorCalendarIsGroupMailboxForTentative string = `ErrorCalendarIsGroupMailboxForTentative`
	// This error indicates that the operation (currently CancelItem) on the calendar item is not valid for an attendee. Only the meeting organizer can cancel the meeting.
	ResponseCodeErrorCalendarIsNotOrganizer string = `ErrorCalendarIsNotOrganizer`
	// This error indicates that AcceptItem is invalid for the organizer's calendar item.
	ResponseCodeErrorCalendarIsOrganizerForAccept string = `ErrorCalendarIsOrganizerForAccept`
	// This error indicates that DeclineItem is invalid for the organizer's calendar item.
	ResponseCodeErrorCalendarIsOrganizerForDecline string = `ErrorCalendarIsOrganizerForDecline`
	// This error indicates that RemoveItem is invalid for the organizer's calendar item. To remove a meeting from the calendar, the organizer must use CancelCalendarItem.
	ResponseCodeErrorCalendarIsOrganizerForRemove string = `ErrorCalendarIsOrganizerForRemove`
	// This error indicates that TentativelyAcceptItem is invalid for the organizer's calendar item.
	ResponseCodeErrorCalendarIsOrganizerForTentative string = `ErrorCalendarIsOrganizerForTentative`
	// This error indicates that a meeting request is out-of-date and cannot be updated.
	ResponseCodeErrorCalendarMeetingRequestIsOutOfDate string = `ErrorCalendarMeetingRequestIsOutOfDate`
	// This error indicates that the occurrence index does not point to an occurrence within the current recurrence. For example, if your recurrence pattern defines a set of three meeting occurrences and you try to access the fifth occurrence, this response code will result.
	ResponseCodeErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange string = `ErrorCalendarOccurrenceIndexIsOutOfRecurrenceRange`
	// This error indicates that any operation on a deleted occurrence (addressed via recurring master ID and occurrence index) is invalid.
	ResponseCodeErrorCalendarOccurrenceIsDeletedFromRecurrence string = `ErrorCalendarOccurrenceIsDeletedFromRecurrence`
	// This error is reported on CreateItem and UpdateItem operations for calendar items or task recurrence properties when the property value is out of range. For example, specifying the fifteenth week of the month will result in this response code.
	ResponseCodeErrorCalendarOutOfRange string = `ErrorCalendarOutOfRange`
	// This error occurs when attempting to invoke the FindItem operation with a SeekToConditionPageItemView for fetching calendar items, which is not supported.
	ResponseCodeErrorCalendarSeekToConditionNotSupported string = `ErrorCalendarSeekToConditionNotSupported`
	// This error occurs when Start to End range for the CalendarView element is more than the maximum allowed, currently 2 years.
	ResponseCodeErrorCalendarViewRangeTooBig string = `ErrorCalendarViewRangeTooBig`
	// This error indicates that the requesting account is not a valid account in the directory database.
	ResponseCodeErrorCallerIsInvalidADAccount string = `ErrorCallerIsInvalidADAccount`
	// Indicates that an attempt was made to archive a calendar contact task folder.
	ResponseCodeErrorCannotArchiveCalendarContactTaskFolderException string = `ErrorCannotArchiveCalendarContactTaskFolderException`
	// Indicates that attempt was made to archive items in the archive mailbox.
	ResponseCodeErrorCannotArchiveItemsInArchiveMailbox string = `ErrorCannotArchiveItemsInArchiveMailbox`
	// Indicates that an attempt was made to archive items in public folders.
	ResponseCodeErrorCannotArchiveItemsInPublicFolders string = `ErrorCannotArchiveItemsInPublicFolders`
	// This error occurs when a calendar item is being created and the SavedItemFolderId attribute refers to a non-calendar folder.
	ResponseCodeErrorCannotCreateCalendarItemInNonCalendarFolder string = `ErrorCannotCreateCalendarItemInNonCalendarFolder`
	// This error occurs when a contact is being created and the SavedItemFolderId attribute refers to a non-contact folder.
	ResponseCodeErrorCannotCreateContactInNonContactFolder string = `ErrorCannotCreateContactInNonContactFolder`
	// This error indicates that a post item cannot be created in a folder other than a mail folder, such as Calendar, Contact, Tasks, Notes, and so on.
	ResponseCodeErrorCannotCreatePostItemInNonMailFolder string = `ErrorCannotCreatePostItemInNonMailFolder`
	// This error occurs when a task is being created and the SavedItemFolderId attribute refers to a non-task folder.
	ResponseCodeErrorCannotCreateTaskInNonTaskFolder string = `ErrorCannotCreateTaskInNonTaskFolder`
	// This error occurs when the item or folder to delete cannot be deleted.
	ResponseCodeErrorCannotDeleteObject string = `ErrorCannotDeleteObject`
	// The DeleteItem operation returns this error when it fails to delete the current occurrence of a recurring task. This can only happen if the AffectedTaskOccurrences attribute has been set to SpecifiedOccurrenceOnly.
	ResponseCodeErrorCannotDeleteTaskOccurrence string = `ErrorCannotDeleteTaskOccurrence`
	// Indicates that an attempt was made to disable a mandatorty extension.
	ResponseCodeErrorCannotDisableMandatoryExtension string = `ErrorCannotDisableMandatoryExtension`
	// This error must be returned when the server cannot empty a folder.
	ResponseCodeErrorCannotEmptyFolder string = `ErrorCannotEmptyFolder`
	// Specifies that the server could not retrieve the external URL for Outlook Web App Options.
	ResponseCodeErrorCannotGetExternalEcpUrl string = `ErrorCannotGetExternalEcpUrl`
	// Indicates that the source folder path could not be retrieved.
	ResponseCodeErrorCannotGetSourceFolderPath string = `ErrorCannotGetSourceFolderPath`
	// The GetAttachment operation returns this error if it cannot retrieve the body of a file attachment.
	ResponseCodeErrorCannotOpenFileAttachment string = `ErrorCannotOpenFileAttachment`
	// This error indicates that the caller tried to set calendar permissions on a non-calendar folder.
	ResponseCodeErrorCannotSetCalendarPermissionOnNonCalendarFolder string = `ErrorCannotSetCalendarPermissionOnNonCalendarFolder`
	// This error indicates that the caller tried to set non-calendar permissions on a calendar folder.
	ResponseCodeErrorCannotSetNonCalendarPermissionOnCalendarFolder string = `ErrorCannotSetNonCalendarPermissionOnCalendarFolder`
	// This error indicates that you cannot set unknown permissions in a permissions set.
	ResponseCodeErrorCannotSetPermissionUnknownEntries string = `ErrorCannotSetPermissionUnknownEntries`
	// Indicates that an attempt was made to specify the search folder as the source folder.
	ResponseCodeErrorCannotSpecifySearchFolderAsSourceFolder string = `ErrorCannotSpecifySearchFolderAsSourceFolder`
	// This error occurs when a request that requires an item identifier is given a folder identifier.
	ResponseCodeErrorCannotUseFolderIdForItemId string = `ErrorCannotUseFolderIdForItemId`
	// This error occurs when a request that requires a folder identifier is given an item identifier.
	ResponseCodeErrorCannotUseItemIdForFolderId string = `ErrorCannotUseItemIdForFolderId`
	// This response code has been replaced by ErrorChangeKeyRequiredForWriteOperations
	ResponseCodeErrorChangeKeyRequired string = `ErrorChangeKeyRequired`
	// This error is returned when the change key for an item is missing or stale. For SendItem, UpdateItem, and UpdateFolder operations, the caller must pass in a correct and current change key for the item. Note that this is the case with UpdateItem even when conflict resolution is set to always overwrite.
	ResponseCodeErrorChangeKeyRequiredForWriteOperations string = `ErrorChangeKeyRequiredForWriteOperations`
	// Specifies that the client was disconnected.
	ResponseCodeErrorClientDisconnected string = `ErrorClientDisconnected`
	// This error is intended for internal use only.
	ResponseCodeErrorClientIntentInvalidStateDefinition string = `ErrorClientIntentInvalidStateDefinition`
	// This error is intended for internal use only.
	ResponseCodeErrorClientIntentNotFound string = `ErrorClientIntentNotFound`
	// This error occurs when Exchange Web Services cannot connect to the mailbox.
	ResponseCodeErrorConnectionFailed string = `ErrorConnectionFailed`
	// This error indicates that the property that was inspected for a Contains filter is not a string type.
	ResponseCodeErrorContainsFilterWrongType string = `ErrorContainsFilterWrongType`
	// The GetItem operation returns this error when Exchange Web Services is unable to retrieve the MIME content for the item requested. The CreateItem operation returns this error when Exchange Web Services is unable to create the item from the supplied MIME content. Usually this is an indication that the item property is corrupted or truncated.
	ResponseCodeErrorContentConversionFailed string = `ErrorContentConversionFailed`
	// This error occurs when a search request is made using the QueryString option and content indexing is not enabled for the target mailbox.
	ResponseCodeErrorContentIndexingNotEnabled string = `ErrorContentIndexingNotEnabled`
	// This error occurs when the data is corrupted and cannot be processed.
	ResponseCodeErrorCorruptData string = `ErrorCorruptData`
	// This error occurs when the caller does not have permission to create the item.
	ResponseCodeErrorCreateItemAccessDenied string = `ErrorCreateItemAccessDenied`
	// This error occurs when one or more of the managed folders that were specified in the CreateManagedFolder operation request failed to be created. Search for each folder to determine which folders were created and which folders do not exist.
	ResponseCodeErrorCreateManagedFolderPartialCompletion string = `ErrorCreateManagedFolderPartialCompletion`
	// This error occurs when the calling account does not have the permissions required to create the subfolder.
	ResponseCodeErrorCreateSubfolderAccessDenied string = `ErrorCreateSubfolderAccessDenied`
	// This error occurs when an attempt is made to move an item or folder from one mailbox to another. If the source mailbox and destination mailbox are different, you will get this error.
	ResponseCodeErrorCrossMailboxMoveCopy string = `ErrorCrossMailboxMoveCopy`
	// This error indicates that the request is not allowed because the Client Access server that should service the request is in a different site.
	ResponseCodeErrorCrossSiteRequest string = `ErrorCrossSiteRequest`
	// This error can occur in the following scenarios:  - An attempt is made to access or write a property on an item and the property value is too large.  - The base64 encoded MIME content length within the request XML exceeds the limit.  - The size of the body of an existing item body exceeds the limit.  - The consumer tries to set an HTML or text body whose length (or combined length in the case of append) exceeds the limit.
	ResponseCodeErrorDataSizeLimitExceeded string = `ErrorDataSizeLimitExceeded`
	// This error occurs when the underlying data provider fails to complete the operation.
	ResponseCodeErrorDataSourceOperation string = `ErrorDataSourceOperation`
	// This error occurs in an AddDelegate operation when the specified user already exists in the list of delegates.
	ResponseCodeErrorDelegateAlreadyExists string = `ErrorDelegateAlreadyExists`
	// This error occurs in an AddDelegate operation when the specified user to be added is the owner of the mailbox.
	ResponseCodeErrorDelegateCannotAddOwner string = `ErrorDelegateCannotAddOwner`
	// This error occurs in a GetDelegate operation when either there is no delegate information on the local FreeBusy message or no Active Directory public delegate (no "public delegate" or no "Send On Behalf" entry in AD DS).
	ResponseCodeErrorDelegateMissingConfiguration string = `ErrorDelegateMissingConfiguration`
	// This error occurs when a specified user cannot be mapped to a user in AD DS.
	ResponseCodeErrorDelegateNoUser string = `ErrorDelegateNoUser`
	// This error occurs in the AddDelegate operation when an added delegate user is not valid.
	ResponseCodeErrorDelegateValidationFailed string = `ErrorDelegateValidationFailed`
	// This error occurs when an attempt is made to delete a distinguished folder.
	ResponseCodeErrorDeleteDistinguishedFolder string = `ErrorDeleteDistinguishedFolder`
	// This response code is not used.
	ResponseCodeErrorDeleteItemsFailed string = `ErrorDeleteItemsFailed`
	// This error is intended for internal use only.
	ResponseCodeErrorDeleteUnifiedMessagingPromptFailed string = `ErrorDeleteUnifiedMessagingPromptFailed`
	// This error is returned when discovery searches are disabled on a tenant or server. This error was introduced in Exchange 2013.
	ResponseCodeErrorDiscoverySearchesDisabled string = `ErrorDiscoverySearchesDisabled`
	// This error indicates that a distinguished user ID is not valid for the operation. DistinguishedUserType should not be present in the request.
	ResponseCodeErrorDistinguishedUserNotSupported string = `ErrorDistinguishedUserNotSupported`
	// This error indicates that a request distribution list member does not exist in the distribution list.
	ResponseCodeErrorDistributionListMemberNotExist string = `ErrorDistributionListMemberNotExist`
	// This error occurs when duplicate folder names are specified within the FolderNames element of the CreateManagedFolder operation request.
	ResponseCodeErrorDuplicateInputFolderNames string = `ErrorDuplicateInputFolderNames`
	// This error is returned when there are duplicate legacy distinguished names in Active Directory Domain Services (AD DS). This error was introduced in Exchange 2013.
	ResponseCodeErrorDuplicateLegacyDistinguishedName string = `ErrorDuplicateLegacyDistinguishedName`
	// This error indicates that there are duplicate SOAP headers.
	ResponseCodeErrorDuplicateSOAPHeader string = `ErrorDuplicateSOAPHeader`
	// This error indicates that a duplicate user ID has been found in a permission set, either Default or Anonymous are set more than once, or there are duplicate SIDs or recipients.
	ResponseCodeErrorDuplicateUserIdsSpecified string = `ErrorDuplicateUserIdsSpecified`
	// This error occurs when a request attempts to create/update the search parameters of a search folder. For example, this can occur when a search folder is created in the mailbox but the search folder is directed to look in another mailbox.
	ResponseCodeErrorEmailAddressMismatch string = `ErrorEmailAddressMismatch`
	// This error occurs when the event that is associated with a watermark is deleted before the event is returned. When this error is returned, the subscription is also deleted.
	ResponseCodeErrorEventNotFound string = `ErrorEventNotFound`
	// This error -iIndicates that there are more concurrent requests against the server than are allowed by a user's policy.
	ResponseCodeErrorExceededConnectionCount string = `ErrorExceededConnectionCount`
	// This error indicates that a search operation call has exceeded the total number of items that can be returned.
	ResponseCodeErrorExceededFindCountLimit string = `ErrorExceededFindCountLimit`
	// This error indicates that a user's throttling policy maximum subscription count has been exceeded.
	ResponseCodeErrorExceededSubscriptionCount string = `ErrorExceededSubscriptionCount`
	// This error occurs if the GetEvents operation is called as a subscription is being deleted because it has expired.
	ResponseCodeErrorExpiredSubscription string = `ErrorExpiredSubscription`
	// This error occurs when the operation to export remote archive mailbox items failed.
	ResponseCodeErrorExportRemoteArchiveItemsFailed string = `ErrorExportRemoteArchiveItemsFailed`
	// Indicates that the extension was not found.
	ResponseCodeErrorExtensionNotFound string = `ErrorExtensionNotFound`
	// This error occurs when the operation to find the remote archive mailbox folder failed.
	ResponseCodeErrorFindRemoteArchiveFolderFailed string = `ErrorFindRemoteArchiveFolderFailed`
	// This error occurs when the folder is corrupted and cannot be saved.
	ResponseCodeErrorFolderCorrupt string = `ErrorFolderCorrupt`
	// This error occurs when an attempt is made to create a folder that has the same name as another folder in the same parent. Duplicate folder names are not allowed.
	ResponseCodeErrorFolderExists string = `ErrorFolderExists`
	// This error indicates that the folder ID that was specified does not correspond to a valid folder, or that the delegate does not have permission to access the folder.
	ResponseCodeErrorFolderNotFound string = `ErrorFolderNotFound`
	// This error indicates that the requested property could not be retrieved. This does not indicate that the property does not exist, but that the property was corrupted in some way so that the retrieval failed.
	ResponseCodeErrorFolderPropertRequestFailed string = `ErrorFolderPropertRequestFailed`
	// This error indicates that the folder could not be created or updated because of an invalid state.
	ResponseCodeErrorFolderSave string = `ErrorFolderSave`
	// This error indicates that the folder could not be created or updated because of an invalid state.
	ResponseCodeErrorFolderSaveFailed string = `ErrorFolderSaveFailed`
	// This error indicates that the folder could not be created or updated because of invalid property values. The response code lists which properties caused the problem.
	ResponseCodeErrorFolderSavePropertyError string = `ErrorFolderSavePropertyError`
	// Specifies that the maximum group member count has been reached for obtaining free/busy information for a distribution list. This error MUST be returned when the maximum group member count has been reached for obtaining free/busy information for a distribution list.
	ResponseCodeErrorFreeBusyDLLimitReached string = `ErrorFreeBusyDLLimitReached`
	// This error is returned when free/busy information cannot be retrieved because of an intervening failure.
	ResponseCodeErrorFreeBusyGenerationFailed string = `ErrorFreeBusyGenerationFailed`
	// This error occurs when the operation to get the remote archive mailbox folder failed.
	ResponseCodeErrorGetRemoteArchiveFolderFailed string = `ErrorGetRemoteArchiveFolderFailed`
	// This error occurs when the operation to get the remote archive mailbox item failed.
	ResponseCodeErrorGetRemoteArchiveItemFailed string = `ErrorGetRemoteArchiveItemFailed`
	// This response code is not used.
	ResponseCodeErrorGetServerSecurityDescriptorFailed string = `ErrorGetServerSecurityDescriptorFailed`
	// This error indicates that a valid VoIP gateway is not available.
	ResponseCodeErrorIPGatewayNotFound string = `ErrorIPGatewayNotFound`
	// This error is returned when new instant messaging (IM) contacts cannot be added because the maximum number of contacts has been reached. This error was introduced in Exchange Server 2013.
	ResponseCodeErrorImContactLimitReached string = `ErrorImContactLimitReached`
	// This error is returned when an attempt is made to add a group display name when an existing group already has the same display name. This error was introduced in Exchange 2013.
	ResponseCodeErrorImGroupDisplayNameAlreadyExists string = `ErrorImGroupDisplayNameAlreadyExists`
	// This error is returned when new IM groups cannot be added because the maximum number of groups has been reached. This error was introduced in Exchange 2013.
	ResponseCodeErrorImGroupLimitReached string = `ErrorImGroupLimitReached`
	// The error is returned in the server-to-server authorization case for Exchange Impersonation when the caller does not have the proper rights to impersonate the specific user in question. This error maps to the ms-Exch-EPI-May-Impersonate extended Active Directory right.
	ResponseCodeErrorImpersonateUserDenied string = `ErrorImpersonateUserDenied`
	// This error is returned in the server-to-server authorization for Exchange Impersonation when the caller does not have the proper rights to impersonate through the Client Access server that they are making the request against. This maps to the ms-Exch-EPI-Impersonation extended Active Directory right.
	ResponseCodeErrorImpersonationDenied string = `ErrorImpersonationDenied`
	// This error indicates that there was an unexpected error when an attempt was made to perform server-to-server authentication. This response code typically indicates either that the service account that is running the Exchange Web Services application pool is configured incorrectly, that Exchange Web Services cannot talk to the directory, or that a trust between forests is not correctly configured.
	ResponseCodeErrorImpersonationFailed string = `ErrorImpersonationFailed`
	// This error MUST be returned if any rule does not validate.
	ResponseCodeErrorInboxRulesValidationError string = `ErrorInboxRulesValidationError`
	// This error indicates that the request was valid for the current Exchange Server version but was invalid for the request server version that was specified.
	ResponseCodeErrorIncorrectSchemaVersion string = `ErrorIncorrectSchemaVersion`
	// This error indicates that each change description in the UpdateItem or UpdateFolder elements must list only one property to update.
	ResponseCodeErrorIncorrectUpdatePropertyCount string = `ErrorIncorrectUpdatePropertyCount`
	// This error occurs when the request contains too many attendees to resolve. By default, the maximum number of attendees to resolve is 100.
	ResponseCodeErrorIndividualMailboxLimitReached string = `ErrorIndividualMailboxLimitReached`
	// This error occurs when the mailbox server is overloaded. Try your request again later.
	ResponseCodeErrorInsufficientResources string = `ErrorInsufficientResources`
	// This error indicates that Exchange Web Services encountered an error that it could not recover from, and a more specific response code that is associated with the error that occurred does not exist.
	ResponseCodeErrorInternalServerError string = `ErrorInternalServerError`
	// This error indicates that an internal server error occurred and that you should try your request again later.
	ResponseCodeErrorInternalServerTransientError string = `ErrorInternalServerTransientError`
	// This error indicates that the level of access that the caller has on the free/busy data is invalid.
	ResponseCodeErrorInvalidAccessLevel string = `ErrorInvalidAccessLevel`
	// This error indicates an error caused by all invalid arguments passed to the GetMessageTrackingReport operation. This error is returned in the following scenarios:  - The user specified in the sending-as parameter does not exist in the directory.  - The user specified in the sending-as parameter is not unique in the directory.  - The sending-as address is empty.  - The sending-as address is not a valid e-mail address.
	ResponseCodeErrorInvalidArgument string = `ErrorInvalidArgument`
	// This error is returned by the GetAttachment operation or the DeleteAttachment operation when an attachment that corresponds to the specified ID is not found.
	ResponseCodeErrorInvalidAttachmentId string = `ErrorInvalidAttachmentId`
	// This error occurs when you try to bind to an existing search folder by using a complex attachment table restriction. Exchange Web Services only supports simple contains filters against the attachment table. If you try to bind to an existing search folder that has a more complex attachment table restriction (a subfilter), Exchange Web Services cannot render the XML for that filter and returns this response code.  Note that you can still call the GetFolder operation on the folder, but do not request the SearchParameters element.
	ResponseCodeErrorInvalidAttachmentSubfilter string = `ErrorInvalidAttachmentSubfilter`
	// This error occurs when you try to bind to an existing search folder by using a complex attachment table restriction. Exchange Web Services only supports simple contains filters against the attachment table.  If you try to bind to an existing search folder that has a more complex attachment table restriction, Exchange Web Services cannot render the XML for that filter. In this case, the attachment subfilter contains a text filter, but it is not looking at the attachment display name.  Note that you can still call the GetFolder operation on the folder, but do not request the SearchParameters element.
	ResponseCodeErrorInvalidAttachmentSubfilterTextFilter string = `ErrorInvalidAttachmentSubfilterTextFilter`
	// This error indicates that the authorization procedure for the requestor failed.
	ResponseCodeErrorInvalidAuthorizationContext string = `ErrorInvalidAuthorizationContext`
	// This error occurs when a consumer passes in a folder or item identifier with a change key that cannot be parsed. For example, this could be invalid base64 content or an empty string.
	ResponseCodeErrorInvalidChangeKey string = `ErrorInvalidChangeKey`
	// This error indicates that a request to get a client access token was not valid. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidClientAccessTokenRequest string = `ErrorInvalidClientAccessTokenRequest`
	// This error indicates that there was an internal error when an attempt was made to resolve the identity of the caller.
	ResponseCodeErrorInvalidClientSecurityContext string = `ErrorInvalidClientSecurityContext`
	// This error is returned when an attempt is made to set the CompleteDate element value of a task to a time in the future. When it is converted to the local time of the Client Access server, the CompleteDate of a task cannot be set to a value that is later than the local time on the Client Access server.
	ResponseCodeErrorInvalidCompleteDate string = `ErrorInvalidCompleteDate`
	// This error indicates that an invalid e-mail address was provided for a contact.
	ResponseCodeErrorInvalidContactEmailAddress string = `ErrorInvalidContactEmailAddress`
	// This error indicates that an invalid e-mail index value was provided for an e-mail entry.
	ResponseCodeErrorInvalidContactEmailIndex string = `ErrorInvalidContactEmailIndex`
	// This error occurs when the credentials that are used to proxy a request to a different directory service forest fail authentication.
	ResponseCodeErrorInvalidCrossForestCredentials string = `ErrorInvalidCrossForestCredentials`
	// This error indicates that the specified folder permissions are invalid.
	ResponseCodeErrorInvalidDelegatePermission string = `ErrorInvalidDelegatePermission`
	// This error indicates that the specified delegate user ID is invalid.
	ResponseCodeErrorInvalidDelegateUserId string = `ErrorInvalidDelegateUserId`
	// This error occurs during Exchange Impersonation when a caller does not specify a UPN, an e-mail address, or a user SID. This will also occur if the caller specifies one or more of those and the values are empty.
	ResponseCodeErrorInvalidExchangeImpersonationHeaderData string = `ErrorInvalidExchangeImpersonationHeaderData`
	// This error occurs when the bitmask that was passed into an Excludes element restriction is unable to be parsed.
	ResponseCodeErrorInvalidExcludesRestriction string = `ErrorInvalidExcludesRestriction`
	// This response code is not used.
	ResponseCodeErrorInvalidExpressionTypeForSubFilter string = `ErrorInvalidExpressionTypeForSubFilter`
	// This error occurs when the following events take place:  - The caller tries to use an extended property that is not supported by Exchange Web Services.  - The caller passes in an invalid combination of attribute values for an extended property. This also occurs if no attributes are passed. Only certain combinations are allowed.
	ResponseCodeErrorInvalidExtendedProperty string = `ErrorInvalidExtendedProperty`
	// This error occurs when the value section of an extended property does not match the type of the property.  For example, setting an extended property that has PropertyType="String" to an array of integers will result in this error. Any string representation that is not coercible into the type that is specified for the extended property will give this error.
	ResponseCodeErrorInvalidExtendedPropertyValue string = `ErrorInvalidExtendedPropertyValue`
	// Specifies that the sharing invitation sender did not create the sharing invitation metadata. This error code MUST be returned if the sharing invitation sender did not create the sharing invitation metadata.
	ResponseCodeErrorInvalidExternalSharingInitiator string = `ErrorInvalidExternalSharingInitiator`
	// This error indicates that a sharing message is not intended for the caller.
	ResponseCodeErrorInvalidExternalSharingSubscriber string = `ErrorInvalidExternalSharingSubscriber`
	// This error indicates that the requestor's organization federation objects are not correctly configured.
	ResponseCodeErrorInvalidFederatedOrganizationId string = `ErrorInvalidFederatedOrganizationId`
	// This error occurs when the folder ID is corrupt.
	ResponseCodeErrorInvalidFolderId string = `ErrorInvalidFolderId`
	// This error indicates that the specified folder type is invalid for the current operation. For example, you cannot create a Search folder in a public folder.
	ResponseCodeErrorInvalidFolderTypeForOperation string = `ErrorInvalidFolderTypeForOperation`
	// This error occurs in fractional paging when the user has specified one of the following:  - A numerator that is greater than the denominator  - A numerator that is less than zero  - A denominator that is less than or equal to zero.
	ResponseCodeErrorInvalidFractionalPagingParameters string = `ErrorInvalidFractionalPagingParameters`
	// This error occurs when the GetUserAvailability operation is called with a FreeBusyViewType of None.
	ResponseCodeErrorInvalidFreeBusyViewType string = `ErrorInvalidFreeBusyViewType`
	// Specifies that the DataType and ShareFolderId element are both present in a request. This error code MUST be returned if the DataType and ShareFolderId element are both present in a request.
	ResponseCodeErrorInvalidGetSharingFolderRequest string = `ErrorInvalidGetSharingFolderRequest`
	// This error indicates that the ID and/or change key is malformed.
	ResponseCodeErrorInvalidId string = `ErrorInvalidId`
	// This error occurs when the caller specifies an Id attribute that is empty.
	ResponseCodeErrorInvalidIdEmpty string = `ErrorInvalidIdEmpty`
	// This error occurs when the caller specifies an Id attribute that is malformed.
	ResponseCodeErrorInvalidIdMalformed string = `ErrorInvalidIdMalformed`
	// This error indicates that a folder or item ID that is using the Exchange 2007 format was specified for a request with a version of Exchange 2007 SP1 or later. You cannot use Exchange 2007 IDs in Exchange 2007 SP1 or later requests. You have to use the ConvertId operation to convert them first.
	ResponseCodeErrorInvalidIdMalformedEwsLegacyIdFormat string = `ErrorInvalidIdMalformedEwsLegacyIdFormat`
	// This error occurs when the caller specifies an Id attribute that is too long.
	ResponseCodeErrorInvalidIdMonikerTooLong string = `ErrorInvalidIdMonikerTooLong`
	// This error is returned whenever an ID that is not an item attachment ID is passed to a Web service method that expects an attachment ID.
	ResponseCodeErrorInvalidIdNotAnItemAttachmentId string = `ErrorInvalidIdNotAnItemAttachmentId`
	// This error occurs when a contact in your mailbox is corrupt.
	ResponseCodeErrorInvalidIdReturnedByResolveNames string = `ErrorInvalidIdReturnedByResolveNames`
	// This error occurs when the caller specifies an Id attribute that is too long.
	ResponseCodeErrorInvalidIdStoreObjectIdTooLong string = `ErrorInvalidIdStoreObjectIdTooLong`
	// This error is returned when the attachment hierarchy on an item exceeds the maximum of 255 levels deep.
	ResponseCodeErrorInvalidIdTooManyAttachmentLevels string = `ErrorInvalidIdTooManyAttachmentLevels`
	// This response code is not used.
	ResponseCodeErrorInvalidIdXml string = `ErrorInvalidIdXml`
	// This error is returned when the specified IM contact identifier does not represent a valid identifier. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidImContactId string = `ErrorInvalidImContactId`
	// This error is returned when the specified IM distribution group SMTP address identifier does not represent a valid identifier. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidImDistributionGroupSmtpAddress string = `ErrorInvalidImDistributionGroupSmtpAddress`
	// This error is returned when the specified IM group identifier does not represent a valid identifier. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidImGroupId string = `ErrorInvalidImGroupId`
	// This error occurs if the offset for indexed paging is negative.
	ResponseCodeErrorInvalidIndexedPagingParameters string = `ErrorInvalidIndexedPagingParameters`
	// This response code is not used.
	ResponseCodeErrorInvalidInternetHeaderChildNodes string = `ErrorInvalidInternetHeaderChildNodes`
	// This error occurs when an attempt is made to use an AcceptItem response object for an item type other than a meeting request or a calendar item, or when an attempt is made to accept a calendar item occurrence that is in the Deleted Items folder.
	ResponseCodeErrorInvalidItemForOperationAcceptItem string = `ErrorInvalidItemForOperationAcceptItem`
	// Indicates that the item was invalid for an ArchiveItem operation.
	ResponseCodeErrorInvalidItemForOperationArchiveItem string = `ErrorInvalidItemForOperationArchiveItem`
	// This error occurs when an attempt is made to use a CancelItem response object on an item type other than a calendar item.
	ResponseCodeErrorInvalidItemForOperationCancelItem string = `ErrorInvalidItemForOperationCancelItem`
	// This error is returned from a CreateItem operation if the request contains an unsupported item type.  Supported items include the following objects:  - Item  - Message  - CalendarItem  - Task  - Contact  Certain types are created as a side effect of doing something else. Meeting messages, for example, are created when you send a calendar item to attendees; they are not explicitly created.
	ResponseCodeErrorInvalidItemForOperationCreateItem string = `ErrorInvalidItemForOperationCreateItem`
	// This error is returned when an attempt is made to create an item attachment of an unsupported type. Supported item types for item attachments include the following objects:  - Item  - Message  - CalendarItem  - Task  - Contact  For example, if you try to create a MeetingMessage attachment, you'll encounter this response code.
	ResponseCodeErrorInvalidItemForOperationCreateItemAttachment string = `ErrorInvalidItemForOperationCreateItemAttachment`
	// This error occurs when an attempt is made to use a DeclineItem response object for an item type other than a meeting request or a calendar item, or when an attempt is made to decline a calendar item occurrence that is in the Deleted Items folder.
	ResponseCodeErrorInvalidItemForOperationDeclineItem string = `ErrorInvalidItemForOperationDeclineItem`
	// This error indicates that the ExpandDL operation is valid only for private distribution lists.
	ResponseCodeErrorInvalidItemForOperationExpandDL string = `ErrorInvalidItemForOperationExpandDL`
	// This error is returned from a RemoveItem response object if the request specifies an item that is not a meeting cancellation item.
	ResponseCodeErrorInvalidItemForOperationRemoveItem string = `ErrorInvalidItemForOperationRemoveItem`
	// This error is returned from a SendItem operation if the request specifies an item that is not a message item.
	ResponseCodeErrorInvalidItemForOperationSendItem string = `ErrorInvalidItemForOperationSendItem`
	// This error occurs when an attempt is made to use TentativelyAcceptItem for an item type other than a meeting request or a calendar item, or when an attempt is made to tentatively accept a calendar item occurrence that is in the Deleted Items folder.
	ResponseCodeErrorInvalidItemForOperationTentative string = `ErrorInvalidItemForOperationTentative`
	// The user doesn't have a valid license.
	ResponseCodeErrorInvalidLicense string = `ErrorInvalidLicense`
	// This error occurs when the item can't be liked. Versions of Exchange starting with build number 15.00.0913.09 include this value.
	ResponseCodeErrorInvalidLikeRequest string = `ErrorInvalidLikeRequest`
	// This error is for internal use only. This error is not returned.
	ResponseCodeErrorInvalidLogonType string = `ErrorInvalidLogonType`
	// This error indicates that the CreateItem operation or the UpdateItem operation failed while creating or updating a personal distribution list.
	ResponseCodeErrorInvalidMailbox string = `ErrorInvalidMailbox`
	// This error occurs when the structure of the managed folder is corrupted and cannot be rendered.
	ResponseCodeErrorInvalidManagedFolderProperty string = `ErrorInvalidManagedFolderProperty`
	// This error occurs when the quota that is set on the managed folder is less than zero, which indicates a corrupted managed folder.
	ResponseCodeErrorInvalidManagedFolderQuota string = `ErrorInvalidManagedFolderQuota`
	// This error occurs when the size that is set on the managed folder is less than zero, which indicates a corrupted managed folder.
	ResponseCodeErrorInvalidManagedFolderSize string = `ErrorInvalidManagedFolderSize`
	// This error is returned if the ManagementRole header in the SOAP header is incorrect. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidManagementRoleHeader string = `ErrorInvalidManagementRoleHeader`
	// This error occurs when the supplied merged free/busy internal value is invalid. The default minimum value is 5 minutes. The default maximum value is 1440 minutes.
	ResponseCodeErrorInvalidMergedFreeBusyInterval string = `ErrorInvalidMergedFreeBusyInterval`
	// This error occurs when the name is invalid for the ResolveNames operation. For example, a zero length string, a single space, a comma, and a dash are all invalid names.
	ResponseCodeErrorInvalidNameForNameResolution string = `ErrorInvalidNameForNameResolution`
	// This error indicates an error in the Network Service account on the Client Access server.
	ResponseCodeErrorInvalidNetworkServiceContext string = `ErrorInvalidNetworkServiceContext`
	// This response code is not used.
	ResponseCodeErrorInvalidOofParameter string = `ErrorInvalidOofParameter`
	// This is a general error that is used when the requested operation is invalid.  For example, you cannot do the following:  - Perform a "Deep" traversal by using the FindFolder operation on a public folder.  - Move or copy the public folder root.  - Delete an associated item by using any mode except "Hard" delete.  - Move or copy an associated item.
	ResponseCodeErrorInvalidOperation string = `ErrorInvalidOperation`
	// This error indicates that a caller requested free/busy information for a user in another organization but the organizational relationship does not have free/busy enabled.
	ResponseCodeErrorInvalidOrganizationRelationshipForFreeBusy string = `ErrorInvalidOrganizationRelationshipForFreeBusy`
	// This error occurs when a consumer passes in a zero or a negative value for the maximum rows to be returned.
	ResponseCodeErrorInvalidPagingMaxRows string = `ErrorInvalidPagingMaxRows`
	// This error occurs when a consumer passes in an invalid parent folder for an operation. For example, this error is returned if you try to create a folder within a search folder.
	ResponseCodeErrorInvalidParentFolder string = `ErrorInvalidParentFolder`
	// This error is returned when an attempt is made to set a task completion percentage to an invalid value. The value must be between 0 and 100 (inclusive).
	ResponseCodeErrorInvalidPercentCompleteValue string = `ErrorInvalidPercentCompleteValue`
	// This error indicates that the permission level is inconsistent with the permission settings.
	ResponseCodeErrorInvalidPermissionSettings string = `ErrorInvalidPermissionSettings`
	// This error indicates that the caller identifier is not valid.
	ResponseCodeErrorInvalidPhoneCallId string = `ErrorInvalidPhoneCallId`
	// This error indicates that the telephone number is not correct or does not fit the dial plan rules.
	ResponseCodeErrorInvalidPhoneNumber string = `ErrorInvalidPhoneNumber`
	// This error is returned if an invalid photo size is requested from the server. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidPhotoSize string = `ErrorInvalidPhotoSize`
	// This error occurs when the property that you are trying to append to does not support appending. The following are the only properties that support appending:  - Recipient collections (ToRecipients, CcRecipients, BccRecipients)  - Attendee collections (RequiredAttendees, OptionalAttendees, Resources)  - Body  - ReplyTo  In addition, this error occurs when a message body is appended if the format specified in the request does not match the format of the item in the store.
	ResponseCodeErrorInvalidPropertyAppend string = `ErrorInvalidPropertyAppend`
	// This error occurs if the delete operation is specified in an UpdateItem operation or UpdateFolder operation call for a property that does not support the delete operation. For example, you cannot delete the ItemId element of the Item object.
	ResponseCodeErrorInvalidPropertyDelete string = `ErrorInvalidPropertyDelete`
	// This error occurs if the consumer passes in one of the flag properties in an Exists filter. For example, this error occurs if the IsRead or IsFromMe flags are specified in the Exists element. The request should use the IsEqualTo element instead for these as they are flags and therefore part of a single property.
	ResponseCodeErrorInvalidPropertyForExists string = `ErrorInvalidPropertyForExists`
	// This error occurs when the property that you are trying to manipulate does not support the operation that is being performed on it.
	ResponseCodeErrorInvalidPropertyForOperation string = `ErrorInvalidPropertyForOperation`
	// This error occurs if a property that is specified in the request is not available for the item type. For example, this error is returned if a property that is only available on calendar items is requested in a GetItem operation call for a message or is updated in an UpdateItem operation call for a message. This occurs in the following operations:  - GetItem operation  - GetFolder operation  - UpdateItem operation  - UpdateFolder operation
	ResponseCodeErrorInvalidPropertyRequest string = `ErrorInvalidPropertyRequest`
	// This error indicates that the property that you are trying to manipulate does not support the operation that is being performed on it. For example, this error is returned if the property that you are trying to set is read-only.
	ResponseCodeErrorInvalidPropertySet string = `ErrorInvalidPropertySet`
	// This error occurs during an UpdateItem operation when a client tries to update certain properties of a message that has already been sent.  For example, the following properties cannot be updated on a sent message:  - IsReadReceiptRequested  - IsDeliveryReceiptRequested
	ResponseCodeErrorInvalidPropertyUpdateSentMessage string = `ErrorInvalidPropertyUpdateSentMessage`
	// This response code is not used.
	ResponseCodeErrorInvalidProxySecurityContext string = `ErrorInvalidProxySecurityContext`
	// This error occurs if you call the GetEvents operation or the Unsubscribe operation by using a push subscription ID. To unsubscribe from a push subscription, you must respond to a push request with an unsubscribe response, or disconnect your Web service and wait for the push notifications to time out.
	ResponseCodeErrorInvalidPullSubscriptionId string = `ErrorInvalidPullSubscriptionId`
	// This error is returned by the Subscribe operation when it creates a "push" subscription and indicates that the URL that is included in the request is invalid. The following conditions must be met for Exchange Web Services to accept the URL:  - String length > 0 and < 2083.  - Protocol is http or https.  - The URL can be parsed by the URI Microsoft .NET Framework class.
	ResponseCodeErrorInvalidPushSubscriptionUrl string = `ErrorInvalidPushSubscriptionUrl`
	// This error indicates that the search folder has a recipient table filter that Exchange Web Services cannot represent. To get around this error, retrieve the folder without requesting the search parameters.
	ResponseCodeErrorInvalidRecipientSubfilter string = `ErrorInvalidRecipientSubfilter`
	// This error indicates that the search folder has a recipient table filter that Exchange Web Services cannot represent. To get around this error, retrieve the folder without requesting the search parameters.
	ResponseCodeErrorInvalidRecipientSubfilterComparison string = `ErrorInvalidRecipientSubfilterComparison`
	// This error indicates that the search folder has a recipient table filter that Exchange Web Services cannot represent. To get around this error, retrieve the folder without requesting the search parameters.
	ResponseCodeErrorInvalidRecipientSubfilterOrder string = `ErrorInvalidRecipientSubfilterOrder`
	// This error indicates that the search folder has a recipient table filter that Exchange Web Services cannot represent. To get around this error, retrieve the folder without requesting the search parameters.
	ResponseCodeErrorInvalidRecipientSubfilterTextFilter string = `ErrorInvalidRecipientSubfilterTextFilter`
	// This error indicates that the recipient collection on your message or the attendee collection on your calendar item is invalid. For example, this error will be returned when an attempt is made to send an item that has no recipients.
	ResponseCodeErrorInvalidRecipients string = `ErrorInvalidRecipients`
	// This error is returned from the CreateItem operation for Forward and Reply response objects in the following scenarios:  - The referenced item identifier is not a Message, a CalendarItem, or a descendant of a Message or CalendarItem.  - The reference item identifier is for a CalendarItem and the organizer is trying to Reply or ReplyAll to himself.  - The message is a draft and Reply or ReplyAll is selected.  - The reference item, for SuppressReadReceipt, is not a Message or a descendant of a Message.
	ResponseCodeErrorInvalidReferenceItem string = `ErrorInvalidReferenceItem`
	// This error occurs when the SOAP request has a SOAP action header, but nothing in the SOAP body. Note that the SOAP Action header is not required as Exchange Web Services can determine the method to call from the local name of the root element in the SOAP body.
	ResponseCodeErrorInvalidRequest string = `ErrorInvalidRequest`
	// This response code is not used.
	ResponseCodeErrorInvalidRestriction string = `ErrorInvalidRestriction`
	// Indicates that the retention tag GUID was invalid.
	ResponseCodeErrorInvalidRetentionTagIdGuid string = `ErrorInvalidRetentionTagIdGuid`
	// This error is returned when an attempt is made to set an implicit tag on the PolicyTag property. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidRetentionTagInheritance string = `ErrorInvalidRetentionTagInheritance`
	// This error is returned when an attempt is made to set a nonexistent or invisible tag on a PolicyTag property. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidRetentionTagInvisible string = `ErrorInvalidRetentionTagInvisible`
	// This error indicates that no retention tags were found for this user. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidRetentionTagNone string = `ErrorInvalidRetentionTagNone`
	// This error is returned when the specified retention tag has an incorrect action associated with it. This error was introduced in Exchange 2013.
	ResponseCodeErrorInvalidRetentionTagTypeMismatch string = `ErrorInvalidRetentionTagTypeMismatch`
	// This error occurs if the routing type that is passed for a RoutingType (EmailAddressType) element is invalid. Typically, the routing type is set to Simple Mail Transfer Protocol (SMTP).
	ResponseCodeErrorInvalidRoutingType string = `ErrorInvalidRoutingType`
	// This error indicates that the SIP name, dial plan, or the phone number are invalid SIP URIs.
	ResponseCodeErrorInvalidSIPUri string = `ErrorInvalidSIPUri`
	// This error occurs if the specified duration end time is not greater than the start time, or if the end time does not occur in the future.
	ResponseCodeErrorInvalidScheduledOofDuration string = `ErrorInvalidScheduledOofDuration`
	// This error indicates that a proxy request that was sent to another server is not able to service the request due to a versioning mismatch.
	ResponseCodeErrorInvalidSchemaVersionForMailboxVersion string = `ErrorInvalidSchemaVersionForMailboxVersion`
	// This error indicates that the Exchange security descriptor on the Calendar folder in the store is corrupted.
	ResponseCodeErrorInvalidSecurityDescriptor string = `ErrorInvalidSecurityDescriptor`
	// This error occurs during an attempt to send an item where the SavedItemFolderId is specified in the request but the SaveItemToFolder property is set to false.
	ResponseCodeErrorInvalidSendItemSaveSettings string = `ErrorInvalidSendItemSaveSettings`
	// This error indicates that the token that was passed in the header is malformed, does not refer to a valid account in the directory, or is missing the primary group ConnectingSID.
	ResponseCodeErrorInvalidSerializedAccessToken string = `ErrorInvalidSerializedAccessToken`
	// This error indicates that an invalid request server version was specified in the request.
	ResponseCodeErrorInvalidServerVersion string = `ErrorInvalidServerVersion`
	// This error indicates that the sharing metadata is not valid. This can be caused by invalid XML.
	ResponseCodeErrorInvalidSharingData string = `ErrorInvalidSharingData`
	// This error indicates that the sharing message is not valid. This can be caused by a missing property.
	ResponseCodeErrorInvalidSharingMessage string = `ErrorInvalidSharingMessage`
	// This error occurs when an invalid SID is passed in a request.
	ResponseCodeErrorInvalidSid string = `ErrorInvalidSid`
	// This error occurs when the SMTP address cannot be parsed.
	ResponseCodeErrorInvalidSmtpAddress string = `ErrorInvalidSmtpAddress`
	// This response code is not used.
	ResponseCodeErrorInvalidSubfilterType string = `ErrorInvalidSubfilterType`
	// This response code is not used.
	ResponseCodeErrorInvalidSubfilterTypeNotAttendeeType string = `ErrorInvalidSubfilterTypeNotAttendeeType`
	// This response code is not used.
	ResponseCodeErrorInvalidSubfilterTypeNotRecipientType string = `ErrorInvalidSubfilterTypeNotRecipientType`
	// This error indicates that the subscription is no longer valid. This could be because the Client Access server is restarting or because the subscription is expired.
	ResponseCodeErrorInvalidSubscription string = `ErrorInvalidSubscription`
	// This error indicates that the subscribe request included multiple public folder IDs. A subscription can include multiple folders from the same mailbox or one public folder ID.
	ResponseCodeErrorInvalidSubscriptionRequest string = `ErrorInvalidSubscriptionRequest`
	// This error is returned by SyncFolderItems or SyncFolderHierarchy if the SyncState data that is passed is invalid. To fix this error, you must resynchronize without the sync state. Make sure that if you are persisting sync state BLOBs, you are not accidentally truncating the BLOB.
	ResponseCodeErrorInvalidSyncStateData string = `ErrorInvalidSyncStateData`
	// This error indicates that the specified time interval is invalid. The start time must be greater than or equal to the end time.
	ResponseCodeErrorInvalidTimeInterval string = `ErrorInvalidTimeInterval`
	// This error indicates that an internally inconsistent UserId was specified for a permissions operation. For example, if a distinguished user ID is specified (Default or Anonymous), this error is returned if you also try to specify a SID, or primary SMTP address or display name for this user.
	ResponseCodeErrorInvalidUserInfo string = `ErrorInvalidUserInfo`
	// This error indicates that the user Out of Office (OOF) settings are invalid because of a missing internal or external reply.
	ResponseCodeErrorInvalidUserOofSettings string = `ErrorInvalidUserOofSettings`
	// This error occurs during Exchange Impersonation. The caller passed in an invalid UPN in the SOAP header that was not accessible in the directory.
	ResponseCodeErrorInvalidUserPrincipalName string = `ErrorInvalidUserPrincipalName`
	// This error occurs when an invalid SID is passed in a request.
	ResponseCodeErrorInvalidUserSid string = `ErrorInvalidUserSid`
	// This response code is not used.
	ResponseCodeErrorInvalidUserSidMissingUPN string = `ErrorInvalidUserSidMissingUPN`
	// This error indicates that the comparison value in the restriction is invalid for the property you are comparing against. For example, the comparison value of DateTimeCreated > true would return this response code. This response code is also returned if you specify an enumeration property in the comparison, but the value that you are comparing against is not a valid value for that enumeration.
	ResponseCodeErrorInvalidValueForProperty string = `ErrorInvalidValueForProperty`
	// This error is caused by an invalid watermark.
	ResponseCodeErrorInvalidWatermark string = `ErrorInvalidWatermark`
	// This error indicates that conflict resolution was unable to resolve changes for the properties. The items in the store may have been changed and have to be updated. Retrieve the updated change key and try again.
	ResponseCodeErrorIrresolvableConflict string = `ErrorIrresolvableConflict`
	// This error indicates that the state of the object is corrupted and cannot be retrieved. When you are retrieving an item, only certain elements will be in this state, such as Body and MimeContent. Omit these elements and retry the operation.
	ResponseCodeErrorItemCorrupt string = `ErrorItemCorrupt`
	// This error occurs when the item was not found or you do not have permission to access the item.
	ResponseCodeErrorItemNotFound string = `ErrorItemNotFound`
	// This error occurs if a property request on an item fails. The property may exist, but it could not be retrieved.
	ResponseCodeErrorItemPropertyRequestFailed string = `ErrorItemPropertyRequestFailed`
	// This error occurs when attempts to save the item or folder fail.
	ResponseCodeErrorItemSave string = `ErrorItemSave`
	// This error occurs when attempts to save the item or folder fail because of invalid property values. The response code includes the path of the invalid properties.
	ResponseCodeErrorItemSavePropertyError string = `ErrorItemSavePropertyError`
	// This response code is not used.
	ResponseCodeErrorLegacyMailboxFreeBusyViewTypeNotMerged string = `ErrorLegacyMailboxFreeBusyViewTypeNotMerged`
	// This response code is not used.
	ResponseCodeErrorLocalServerObjectNotFound string = `ErrorLocalServerObjectNotFound`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorLocationServicesDisabled string = `ErrorLocationServicesDisabled`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorLocationServicesInvalidRequest string = `ErrorLocationServicesInvalidRequest`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorLocationServicesRequestFailed string = `ErrorLocationServicesRequestFailed`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorLocationServicesRequestTimedOut string = `ErrorLocationServicesRequestTimedOut`
	// This error indicates that the Availability service was unable to log on as the network service to proxy requests to the appropriate sites or forests. This response typically indicates a configuration error.
	ResponseCodeErrorLogonAsNetworkServiceFailed string = `ErrorLogonAsNetworkServiceFailed`
	// This error occurs if the MailboxData element information cannot be mapped to a valid mailbox account.
	ResponseCodeErrorMailRecipientNotFound string = `ErrorMailRecipientNotFound`
	// This error indicates that mail tips are disabled.
	ResponseCodeErrorMailTipsDisabled string = `ErrorMailTipsDisabled`
	// This error indicates that the mailbox information in AD DS is configured incorrectly.
	ResponseCodeErrorMailboxConfiguration string = `ErrorMailboxConfiguration`
	// This error indicates that the MailboxDataArray element in the request is empty. You must supply at least one mailbox identifier.
	ResponseCodeErrorMailboxDataArrayEmpty string = `ErrorMailboxDataArrayEmpty`
	// This error occurs when more than 100 entries are supplied in a MailboxDataArray element..
	ResponseCodeErrorMailboxDataArrayTooBig string = `ErrorMailboxDataArrayTooBig`
	// This error indicates that an attempt to access a mailbox failed because the mailbox is in a failover process.
	ResponseCodeErrorMailboxFailover string = `ErrorMailboxFailover`
	// Indicates that the mailbox hold was not found.
	ResponseCodeErrorMailboxHoldNotFound string = `ErrorMailboxHoldNotFound`
	// This error occurs when the connection to the mailbox to get the calendar view information failed.
	ResponseCodeErrorMailboxLogonFailed string = `ErrorMailboxLogonFailed`
	// This error indicates that the mailbox is being moved to a different mailbox store or server. This error can also indicate that the mailbox is on another server or mailbox database.
	ResponseCodeErrorMailboxMoveInProgress string = `ErrorMailboxMoveInProgress`
	// This error is returned when a scoped search attempt is performed without using a QueryString (String) element for a content indexing search. This is applicable to the SearchMailboxes and FindConversation operations. This error was introduced in Exchange 2013.
	ResponseCodeErrorMailboxScopeNotAllowedWithoutQueryString string = `ErrorMailboxScopeNotAllowedWithoutQueryString`
	// This error indicates that one of the following error conditions occurred:  - The mailbox store is corrupt.  - The mailbox store is being stopped.  - The mailbox store is offline.  - A network error occurred when an attempt was made to access the mailbox store.  - The mailbox store is overloaded and cannot accept any more connections.  - The mailbox store has been paused.
	ResponseCodeErrorMailboxStoreUnavailable string = `ErrorMailboxStoreUnavailable`
	// This error occurs if the managed folder that you are trying to create already exists in a mailbox.
	ResponseCodeErrorManagedFolderAlreadyExists string = `ErrorManagedFolderAlreadyExists`
	// This error occurs when the folder name that was specified in the request does not map to a managed folder definition in AD DS. You can only create instances of managed folders for folders that are defined in AD DS. Check the name and try again.
	ResponseCodeErrorManagedFolderNotFound string = `ErrorManagedFolderNotFound`
	// This error indicates that the managed folders root was deleted from the mailbox or that a folder exists in the same parent folder that has the name of the managed folder root. This will also occur if the attempt to create the root managed folder fails.
	ResponseCodeErrorManagedFoldersRootFailure string = `ErrorManagedFoldersRootFailure`
	// This error indicates that the suggestions engine encountered a problem when it was trying to generate the suggestions.
	ResponseCodeErrorMeetingSuggestionGenerationFailed string = `ErrorMeetingSuggestionGenerationFailed`
	// This error occurs if the MessageDisposition attribute is not set. This attribute is required for the following:  - The CreateItem operation and the UpdateItem operation when the item being created or updated is a Message.  - CancelCalendarItem, AcceptItem, DeclineItem, or TentativelyAcceptItem response objects.
	ResponseCodeErrorMessageDispositionRequired string = `ErrorMessageDispositionRequired`
	// The message per folder receive quota has been exceeded.
	ResponseCodeErrorMessagePerFolderCountReceiveQuotaExceeded string = `ErrorMessagePerFolderCountReceiveQuotaExceeded`
	// This error indicates that the message that you are trying to send exceeds the allowed limits.
	ResponseCodeErrorMessageSizeExceeded string = `ErrorMessageSizeExceeded`
	// This error indicates that the given domain cannot be found.
	ResponseCodeErrorMessageTrackingNoSuchDomain string = `ErrorMessageTrackingNoSuchDomain`
	// This error indicates that the message tracking service cannot track the message.
	ResponseCodeErrorMessageTrackingPermanentError string = `ErrorMessageTrackingPermanentError`
	// This error indicates that the message tracking service is either down or busy. This error code indicates a transient error. Clients can retry to connect to the server when this error is received.
	ResponseCodeErrorMessageTrackingTransientError string = `ErrorMessageTrackingTransientError`
	// This error occurs when the MIME content is not a valid iCal for a CreateItem operation. For a GetItem operation, this response indicates that the MIME content could not be generated.
	ResponseCodeErrorMimeContentConversionFailed string = `ErrorMimeContentConversionFailed`
	// This error occurs when the MIME content is invalid.
	ResponseCodeErrorMimeContentInvalid string = `ErrorMimeContentInvalid`
	// This error occurs when the MIME content in the request is not a valid base 64 string.
	ResponseCodeErrorMimeContentInvalidBase6FourString string = `ErrorMimeContentInvalidBase64String`
	// This error MUST be returned when event notifications are missed.
	ResponseCodeErrorMissedNotificationEvents string = `ErrorMissedNotificationEvents`
	// This error indicates that a required argument was missing from the request. The response message text indicates which argument to check.
	ResponseCodeErrorMissingArgument string = `ErrorMissingArgument`
	// This error indicates that you specified a distinguished folder ID in the request, but the account that made the request does not have a mailbox on the system. In that case, you must supply a Mailbox sub-element under DistinguishedFolderId.
	ResponseCodeErrorMissingEmailAddress string = `ErrorMissingEmailAddress`
	// This error indicates that you specified a distinguished folder ID in the request, but the account that made the request does not have a mailbox on the system. In that case, you must supply a Mailbox sub-element under DistinguishedFolderId. This response is returned from the CreateManagedFolder operation.
	ResponseCodeErrorMissingEmailAddressForManagedFolder string = `ErrorMissingEmailAddressForManagedFolder`
	// This error occurs if the EmailAddress (NonEmptyStringType) element is missing.
	ResponseCodeErrorMissingInformationEmailAddress string = `ErrorMissingInformationEmailAddress`
	// This error occurs if the ReferenceItemId is missing.
	ResponseCodeErrorMissingInformationReferenceItemId string = `ErrorMissingInformationReferenceItemId`
	// This error code is never returned.
	ResponseCodeErrorMissingInformationSharingFolderId string = `ErrorMissingInformationSharingFolderId`
	// This error is returned when an attempt is made to not include the item element in the ItemAttachment element of a CreateAttachment operation request.
	ResponseCodeErrorMissingItemForCreateItemAttachment string = `ErrorMissingItemForCreateItemAttachment`
	// This error occurs when the policy IDs property, property tag 0x6732, for the folder is missing. You should consider this a corrupted folder.
	ResponseCodeErrorMissingManagedFolderId string = `ErrorMissingManagedFolderId`
	// This error indicates that you tried to send an item without including recipients. Note that if you call the CreateItem operation with a message disposition that causes the message to be sent, you will get the following response code: ErrorInvalidRecipients.
	ResponseCodeErrorMissingRecipients string = `ErrorMissingRecipients`
	// This error indicates that a UserId has not been fully specified in a permissions set.
	ResponseCodeErrorMissingUserIdInformation string = `ErrorMissingUserIdInformation`
	// This error indicates that you have specified more than one ExchangeImpersonation element value within a request.
	ResponseCodeErrorMoreThanOneAccessModeSpecified string = `ErrorMoreThanOneAccessModeSpecified`
	// This error indicates that the move or copy operation failed. Moving occurs in the CreateItem operation when you accept a meeting request that is in the Deleted Items folder. In addition, if you decline a meeting request, cancel a calendar item, or remove a meeting from your calendar, it is moved to the Deleted Items folder.
	ResponseCodeErrorMoveCopyFailed string = `ErrorMoveCopyFailed`
	// This error occurs if you try to move a distinguished folder.
	ResponseCodeErrorMoveDistinguishedFolder string = `ErrorMoveDistinguishedFolder`
	// This error occurs when a request attempts to access multiple mailbox servers. This error was introduced in Exchange 2013.
	ResponseCodeErrorMultiLegacyMailboxAccess string = `ErrorMultiLegacyMailboxAccess`
	// This error occurs if the ResolveNames operation returns more than one result or the ambiguous name that you specified matches more than one object in the directory. The response code includes the matched names in the response data.
	ResponseCodeErrorNameResolutionMultipleResults string = `ErrorNameResolutionMultipleResults`
	// This error indicates that the caller does not have a mailbox on the system. The ResolveNames operation or ExpandDL operation is invalid for connecting a user without a mailbox.
	ResponseCodeErrorNameResolutionNoMailbox string = `ErrorNameResolutionNoMailbox`
	// This error indicates that the ResolveNames operation returns no results.
	ResponseCodeErrorNameResolutionNoResults string = `ErrorNameResolutionNoResults`
	// This error MUST be returned to the first subscription connection if a second subscription connection is opened.
	ResponseCodeErrorNewEventStreamConnectionOpened string = `ErrorNewEventStreamConnectionOpened`
	// This error code MUST be returned when the Web service cannot find a server to handle the request.
	ResponseCodeErrorNoApplicableProxyCASServersAvailable string = `ErrorNoApplicableProxyCASServersAvailable`
	// This error occurs if there is no Calendar folder for the mailbox.
	ResponseCodeErrorNoCalendar string = `ErrorNoCalendar`
	// This error indicates that the request referred to a mailbox in another Active Directory site, but no Client Access servers in the destination site were configured for Windows Authentication, and therefore the request could not be proxied.
	ResponseCodeErrorNoDestinationCASDueToKerberosRequirements string = `ErrorNoDestinationCASDueToKerberosRequirements`
	// This error indicates that the request referred to a mailbox in another Active Directory site, but no Client Access servers in the destination site were configured for SSL connections, and therefore the request could not be proxied.
	ResponseCodeErrorNoDestinationCASDueToSSLRequirements string = `ErrorNoDestinationCASDueToSSLRequirements`
	// This error indicates that the request referred to a mailbox in another Active Directory site, but no Client Access servers in the destination site were of an acceptable product version to receive the request, and therefore the request could not be proxied.
	ResponseCodeErrorNoDestinationCASDueToVersionMismatch string = `ErrorNoDestinationCASDueToVersionMismatch`
	// This error occurs if you set the FolderClass element when you are creating an item other than a generic folder. For typed folders such as CalendarFolder and TasksFolder, the folder class is implied. Setting the folder class to a different folder type by using the UpdateFolder operation results in the ErrorObjectTypeChanged response. Instead, use a generic folder type but set the folder class to the value that you require. Exchange Web Services will create the correct strongly typed folder.
	ResponseCodeErrorNoFolderClassOverride string = `ErrorNoFolderClassOverride`
	// This error indicates that the caller does not have free/busy viewing rights on the Calendar folder in question.
	ResponseCodeErrorNoFreeBusyAccess string = `ErrorNoFreeBusyAccess`
	// This error indicates that MAPI properties in the custom range, 0x8000 and greater, cannot be referenced by property tags. You must use the EWS Managed API PropertySetIdproperty or the EWS ExtendedFieldURI element with the PropertySetId attribute.
	ResponseCodeErrorNoPropertyTagForCustomProperties string = `ErrorNoPropertyTagForCustomProperties`
	// This response code is not used.
	ResponseCodeErrorNoPublicFolderReplicaAvailable string = `ErrorNoPublicFolderReplicaAvailable`
	// This error code MUST be returned if no public folder server is available or if the caller does not have a home public server.
	ResponseCodeErrorNoPublicFolderServerAvailable string = `ErrorNoPublicFolderServerAvailable`
	// This error indicates that the request referred to a mailbox in another Active Directory site, but none of the Client Access servers in that site responded, and therefore the request could not be proxied.
	ResponseCodeErrorNoRespondingCASInDestinationSite string = `ErrorNoRespondingCASInDestinationSite`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorNoSpeechDetected string = `ErrorNoSpeechDetected`
	// This error occurs in the following scenarios:  - The e-mail address is empty in CreateManagedFolder.  - The e-mail address does not refer to a valid account in a request that takes an e-mail address in the body or in the SOAP header, such as in an Exchange Impersonation call.
	ResponseCodeErrorNonExistentMailbox string = `ErrorNonExistentMailbox`
	// This error occurs when a caller passes in a non-primary SMTP address. The response includes the correct SMTP address to use.
	ResponseCodeErrorNonPrimarySmtpAddress string = `ErrorNonPrimarySmtpAddress`
	// Specifies that the caller attempted to grant permissions in its calendar or contacts folder to a user in another organization and the attempt failed. This error code MUST be returned when the sharing policy is disabled for the caller or when the sharing policy assigned to the caller disallows sharing for the requested level or the requested folder type.
	ResponseCodeErrorNotAllowedExternalSharingByPolicy string = `ErrorNotAllowedExternalSharingByPolicy`
	// This error indicates that the user is not a delegate for the mailbox. It is returned by the GetDelegate operation, the RemoveDelegate operation, and the UpdateDelegate operation when the specified delegate user is not found in the list of delegates.
	ResponseCodeErrorNotDelegate string = `ErrorNotDelegate`
	// This error indicates that the operation could not be completed because of insufficient memory.
	ResponseCodeErrorNotEnoughMemory string = `ErrorNotEnoughMemory`
	// This error indicates that the sharing message is not supported.
	ResponseCodeErrorNotSupportedSharingMessage string = `ErrorNotSupportedSharingMessage`
	// This error occurs if the object type changed.
	ResponseCodeErrorObjectTypeChanged string = `ErrorObjectTypeChanged`
	// This error occurs when the Start or End time of an occurrence is updated so that the occurrence is scheduled to happen earlier or later than the corresponding previous or next occurrence.
	ResponseCodeErrorOccurrenceCrossingBoundary string = `ErrorOccurrenceCrossingBoundary`
	// This error indicates that the time allotment for a given occurrence overlaps with another occurrence of the same recurring item. This response also occurs when the length in minutes of a given occurrence is larger than Int32.MaxValue.
	ResponseCodeErrorOccurrenceTimeSpanTooBig string = `ErrorOccurrenceTimeSpanTooBig`
	// This error indicates that the current operation is not valid for the public folder root.
	ResponseCodeErrorOperationNotAllowedWithPublicFolderRoot string = `ErrorOperationNotAllowedWithPublicFolderRoot`
	// The tenant is marked for removal.
	ResponseCodeErrorOrganizationAccessBlocked string = `ErrorOrganizationAccessBlocked`
	// Specifies that the requestor's organization is not federated so the requestor cannot create sharing messages to send to an external user or cannot accept sharing messages received from an external user. This error code MUST be returned if the requestor's organization is not federated.
	ResponseCodeErrorOrganizationNotFederated string = `ErrorOrganizationNotFederated`
	// This error MUST be returned when an attempt to manage Inbox rules occurs after another client has accessed the Inbox rules.
	ResponseCodeErrorOutlookRuleBlobExists string = `ErrorOutlookRuleBlobExists`
	// This response code is not used.
	ResponseCodeErrorParentFolderIdRequired string = `ErrorParentFolderIdRequired`
	// This error occurs in the CreateFolder operation when the parent folder is not found.
	ResponseCodeErrorParentFolderNotFound string = `ErrorParentFolderNotFound`
	// This error indicates that you must change your password before you can access this mailbox. This occurs when a new account has been created and the administrator indicated that the user must change the password at first logon. You cannot update the password by using Exchange Web Services. You must use a tool such as Microsoft Office Outlook Web App to change your password.
	ResponseCodeErrorPasswordChangeRequired string = `ErrorPasswordChangeRequired`
	// This error indicates that the password has expired. You cannot change the password by using Exchange Web Services. You must use a tool such as Outlook Web App to change your password.
	ResponseCodeErrorPasswordExpired string = `ErrorPasswordExpired`
	// This error indicates that the requester tried to grant permissions in its calendar or contacts folder to an external user but the sharing policy assigned to the requester indicates that the requested permission level is higher than what the sharing policy allows.
	ResponseCodeErrorPermissionNotAllowedByPolicy string = `ErrorPermissionNotAllowedByPolicy`
	// This error indicates that the telephone number was not in the correct form.
	ResponseCodeErrorPhoneNumberNotDialable string = `ErrorPhoneNumberNotDialable`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorPromptPublishingOperationFailed string = `ErrorPromptPublishingOperationFailed`
	// This error indicates that the update failed because of invalid property values. The response message includes the invalid property paths.
	ResponseCodeErrorPropertyUpdate string = `ErrorPropertyUpdate`
	// This response code is not used.
	ResponseCodeErrorPropertyValidationFailure string = `ErrorPropertyValidationFailure`
	// This error indicates that the request referred to a subscription that exists on another Client Access server, but an attempt to proxy the request to that Client Access server failed.
	ResponseCodeErrorProxiedSubscriptionCallFailure string = `ErrorProxiedSubscriptionCallFailure`
	// This response code is not used.
	ResponseCodeErrorProxyCallFailed string = `ErrorProxyCallFailed`
	// This error indicates that the request referred to a mailbox in another Active Directory site, and the original caller is a member of more than 3,000 groups.
	ResponseCodeErrorProxyGroupSidLimitExceeded string = `ErrorProxyGroupSidLimitExceeded`
	// This error indicates that the request that Exchange Web Services sent to another Client Access server when trying to fulfill a GetUserAvailabilityRequest request was invalid. This response code typically indicates that a configuration or rights error has occurred, or that someone tried unsuccessfully to mimic an availability proxy request.
	ResponseCodeErrorProxyRequestNotAllowed string = `ErrorProxyRequestNotAllowed`
	// This error indicates that Exchange Web Services tried to proxy an availability request to another Client Access server for fulfillment, but the request failed. This response can be caused by network connectivity issues or request timeout issues.
	ResponseCodeErrorProxyRequestProcessingFailed string = `ErrorProxyRequestProcessingFailed`
	// This error code must be returned if the Web service cannot determine whether the request is to run on the target server or will be proxied to another server.
	ResponseCodeErrorProxyServiceDiscoveryFailed string = `ErrorProxyServiceDiscoveryFailed`
	// This response code is not used.
	ResponseCodeErrorProxyTokenExpired string = `ErrorProxyTokenExpired`
	// This error occurs when the public folder mailbox URL cannot be found. This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorPublicFolderMailboxDiscoveryFailed string = `ErrorPublicFolderMailboxDiscoveryFailed`
	// This error occurs when an attempt is made to access a public folder and the attempt is unsuccessful. This error was introduced in Exchange 2013Exchange Server 2013.
	ResponseCodeErrorPublicFolderOperationFailed string = `ErrorPublicFolderOperationFailed`
	// This error occurs when the recipient that was passed to the GetUserAvailability operation is located on a computer that is running a version of Exchange Server that is earlier than Exchange 2007, and the request to retrieve free/busy information for the recipient from the public folder server failed.
	ResponseCodeErrorPublicFolderRequestProcessingFailed string = `ErrorPublicFolderRequestProcessingFailed`
	// This error occurs when the recipient that was passed to the GetUserAvailability operation is located on a computer that is running a version of Exchange Server that is earlier than Exchange 2007, and the request to retrieve free/busy information for the recipient from the public folder server failed because the organizational unit did not have an associated public folder server.
	ResponseCodeErrorPublicFolderServerNotFound string = `ErrorPublicFolderServerNotFound`
	// This error occurs when a synchronization operation succeeds against the primary public folder mailbox but does not succeed against the secondary public folder mailbox. This error was introduced in Exchange 2013.
	ResponseCodeErrorPublicFolderSyncException string = `ErrorPublicFolderSyncException`
	// This error indicates that the search folder restriction may be valid, but it is not supported by EWS. Exchange Web Services limits restrictions to contain a maximum of 255 filter expressions. If you try to bind to an existing search folder that exceeds 255, this response code is returned.
	ResponseCodeErrorQueryFilterTooLong string = `ErrorQueryFilterTooLong`
	// This error occurs when the mailbox quota is exceeded.
	ResponseCodeErrorQuotaExceeded string = `ErrorQuotaExceeded`
	// This error is returned by the GetEvents operation or push notifications when a failure occurs while retrieving event information. When this error is returned, the subscription is deleted. Re-create the event synchronization based on a last known watermark.
	ResponseCodeErrorReadEventsFailed string = `ErrorReadEventsFailed`
	// This error is returned by the CreateItem operation if an attempt was made to suppress a read receipt when the message sender did not request a read receipt on the message or if the message is in the Junk E-mail folder.
	ResponseCodeErrorReadReceiptNotPending string = `ErrorReadReceiptNotPending`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorRecipientNotFound string = `ErrorRecipientNotFound`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorRecognizerNotInstalled string = `ErrorRecognizerNotInstalled`
	// This error occurs when the end date for the recurrence is after 9/1/4500.
	ResponseCodeErrorRecurrenceEndDateTooBig string = `ErrorRecurrenceEndDateTooBig`
	// This error occurs when the specified recurrence does not have any occurrence instances in the specified range.
	ResponseCodeErrorRecurrenceHasNoOccurrence string = `ErrorRecurrenceHasNoOccurrence`
	// This error indicates that the delegate list failed to be saved after delegates were removed.
	ResponseCodeErrorRemoveDelegatesFailed string = `ErrorRemoveDelegatesFailed`
	// This response code is not used.
	ResponseCodeErrorRequestAborted string = `ErrorRequestAborted`
	// This error occurs when the request stream is larger than 400 KB.
	ResponseCodeErrorRequestStreamTooBig string = `ErrorRequestStreamTooBig`
	// This error is returned when a required property is missing in a CreateAttachment operation request. The missing property URI is included in the response.
	ResponseCodeErrorRequiredPropertyMissing string = `ErrorRequiredPropertyMissing`
	// This error indicates that the caller has specified a folder that is not a contacts folder to the ResolveNames operation.
	ResponseCodeErrorResolveNamesInvalidFolderType string = `ErrorResolveNamesInvalidFolderType`
	// This error indicates that the caller has specified more than one contacts folder to the ResolveNames operation.
	ResponseCodeErrorResolveNamesOnlyOneContactsFolderAllowed string = `ErrorResolveNamesOnlyOneContactsFolderAllowed`
	// This response code is not used.
	ResponseCodeErrorResponseSchemaValidation string = `ErrorResponseSchemaValidation`
	// This error occurs when the restriction cannot be evaluated by Exchange Web Services.
	ResponseCodeErrorRestrictionTooComplex string = `ErrorRestrictionTooComplex`
	// This error occurs if the restriction contains more than 255 nodes.
	ResponseCodeErrorRestrictionTooLong string = `ErrorRestrictionTooLong`
	// This error indicates that the number of calendar entries for a given recipient exceeds the allowed limit of 1000. Reduce the window and try again.
	ResponseCodeErrorResultSetTooBig string = `ErrorResultSetTooBig`
	// This error MUST be returned when a user's rule quota has been exceeded.
	ResponseCodeErrorRulesOverQuota string = `ErrorRulesOverQuota`
	// This error occurs when the SavedItemFolderId is not found.
	ResponseCodeErrorSavedItemFolderNotFound string = `ErrorSavedItemFolderNotFound`
	// This error occurs when the request cannot be validated against the schema.
	ResponseCodeErrorSchemaValidation string = `ErrorSchemaValidation`
	// This error indicates that the search folder was created, but the search criteria were never set on the folder. This only occurs when you access corrupted search folders that were created by using another API or client. To fix this error, use the UpdateFolder operation to set the SearchParameters element to include the restriction that should be on the folder.
	ResponseCodeErrorSearchFolderNotInitialized string = `ErrorSearchFolderNotInitialized`
	// This error is returned when an unexpected photo size is requested in a GetUserPhoto operation request. This error was introduced in Exchange 2013.
	ResponseCodeErrorSearchQueryHasTooManyKeywords string = `ErrorSearchQueryHasTooManyKeywords`
	// This error is returned when a SearchMailboxes operation request contains too many mailboxes to search. This error was introduced in Exchange 2013.
	ResponseCodeErrorSearchTooManyMailboxes string = `ErrorSearchTooManyMailboxes`
	// This error occurs when both of the following conditions occur:  - A user has been granted CanActAsOwner permissions but is not granted delegate rights on the principal's mailbox.  - The same user tries to create and send an e-mail message in the principal's mailbox by using the SendAndSaveCopy option.  The result is an ErrorSendAsDenied error and the creation of the e-mail message in the principal's Drafts folder.
	ResponseCodeErrorSendAsDenied string = `ErrorSendAsDenied`
	// This error is returned by the DeleteItem operation if the SendMeetingCancellations attribute is missing from the request and the item to delete is a calendar item.
	ResponseCodeErrorSendMeetingCancellationsRequired string = `ErrorSendMeetingCancellationsRequired`
	// This error is returned by the UpdateItem operation if the SendMeetingInvitationsOrCancellations attribute is missing from the request and the item to update is a calendar item.
	ResponseCodeErrorSendMeetingInvitationsOrCancellationsRequired string = `ErrorSendMeetingInvitationsOrCancellationsRequired`
	// This error is returned by the CreateItem operation if the SendMeetingInvitations attribute is missing from the request and the item to create is a calendar item.
	ResponseCodeErrorSendMeetingInvitationsRequired string = `ErrorSendMeetingInvitationsRequired`
	// This error indicates that after the organizer sends a meeting request, the request cannot be updated. To modify the meeting, modify the calendar item, not the meeting request.
	ResponseCodeErrorSentMeetingRequestUpdate string = `ErrorSentMeetingRequestUpdate`
	// This error indicates that after the task initiator sends a task request, that request cannot be updated.
	ResponseCodeErrorSentTaskRequestUpdate string = `ErrorSentTaskRequestUpdate`
	// This error occurs when the server is busy.
	ResponseCodeErrorServerBusy string = `ErrorServerBusy`
	// This error indicates that Exchange Web Services tried to proxy a user availability request to the appropriate forest for the recipient, but it could not determine where to send the request because of a service discovery failure.
	ResponseCodeErrorServiceDiscoveryFailed string = `ErrorServiceDiscoveryFailed`
	// Specifies that the external URL property has not been set in the Active Directory database. This error code MUST be returned if the external URL property has not been set in the Active Directory database.
	ResponseCodeErrorSharingNoExternalEwsAvailable string = `ErrorSharingNoExternalEwsAvailable`
	// This error indicates that an attempt at synchronizing a sharing folder failed. This error code is returned when the following occurs:  - The subscription for a sharing folder is not found.  - The sharing folder is not found.  - The corresponding directory user is not found.  - The user no longer exists.  - The appointment is invalid.  - The contact item is invalid.  - There is a communication failure with the remote server.
	ResponseCodeErrorSharingSynchronizationFailed string = `ErrorSharingSynchronizationFailed`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorSpeechGrammarError string = `ErrorSpeechGrammarError`
	// This error occurs in an UpdateItem operation or a SendItem operation when the change key is not up-to-date or was not supplied. Call the GetItem operation to retrieve an updated change key and then try the operation again.
	ResponseCodeErrorStaleObject string = `ErrorStaleObject`
	// This error Indicates that a user cannot immediately send more requests because the submission quota has been reached.
	ResponseCodeErrorSubmissionQuotaExceeded string = `ErrorSubmissionQuotaExceeded`
	// This error occurs when you try to access a subscription by using an account that did not create that subscription. Each subscription can only be accessed by the creator of the subscription.
	ResponseCodeErrorSubscriptionAccessDenied string = `ErrorSubscriptionAccessDenied`
	// This error indicates that you cannot create a subscription if you are not the owner or do not have owner access to the mailbox.
	ResponseCodeErrorSubscriptionDelegateAccessNotSupported string = `ErrorSubscriptionDelegateAccessNotSupported`
	// This error occurs if the subscription that corresponds to the specified SubscriptionId (GetEvents) is not found. The subscription may have expired, the Exchange Web Services process may have been restarted, or an invalid subscription was passed in. If the subscription was valid, re-create the subscription with the latest watermark. This is returned by the Unsubscribe operation or the GetEvents operation responses.
	ResponseCodeErrorSubscriptionNotFound string = `ErrorSubscriptionNotFound`
	// This error code must be returned when a request is made for a subscription that has been unsubscribed.
	ResponseCodeErrorSubscriptionUnsubsribed string = `ErrorSubscriptionUnsubsribed`
	// This error is returned by the SyncFolderItems operation if the parent folder that is specified cannot be found.
	ResponseCodeErrorSyncFolderNotFound string = `ErrorSyncFolderNotFound`
	// This error code is not used. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxActiveToPendingDelete string = `ErrorTeamMailboxActiveToPendingDelete`
	// This error indicates a general error that can occur when trying to access a team mailbox. Try submitting the request at a later time. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxErrorUnknown string = `ErrorTeamMailboxErrorUnknown`
	// This error indicates that an attempt to send a notification to the team mailbox owners was unsuccessful. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxFailedSendingNotifications string = `ErrorTeamMailboxFailedSendingNotifications`
	// This error code is not used. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxNotAuthorizedOwner string = `ErrorTeamMailboxNotAuthorizedOwner`
	// This error indicates that a team mailbox was not found. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxNotFound string = `ErrorTeamMailboxNotFound`
	// This error indicates that a team mailbox was found but that it is not linked to a SharePoint Server. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxNotLinkedToSharePoint string = `ErrorTeamMailboxNotLinkedToSharePoint`
	// This error indicates that a team mailbox was found but that the link to the SharePoint Server is not valid. This error was introduced in Exchange 2013.
	ResponseCodeErrorTeamMailboxUrlValidationFailed string = `ErrorTeamMailboxUrlValidationFailed`
	// This error indicates that the time window that was specified is larger than the allowed limit. By default, the allowed limit is 42.
	ResponseCodeErrorTimeIntervalTooBig string = `ErrorTimeIntervalTooBig`
	// This error indicates that there is a time zone error.
	ResponseCodeErrorTimeZone string = `ErrorTimeZone`
	// This error occurs when there is not enough time to complete the processing of the request.
	ResponseCodeErrorTimeoutExpired string = `ErrorTimeoutExpired`
	// This error indicates that the destination folder does not exist.
	ResponseCodeErrorToFolderNotFound string = `ErrorToFolderNotFound`
	// This error occurs if the caller tries to do a Token serialization request but does not have the ms-Exch-EPI-TokenSerialization right on the Client Access server.
	ResponseCodeErrorTokenSerializationDenied string = `ErrorTokenSerializationDenied`
	// This error occurs when the internal limit on open objects has been exceeded.
	ResponseCodeErrorTooManyObjectsOpened string = `ErrorTooManyObjectsOpened`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorUMServerUnavailable string = `ErrorUMServerUnavailable`
	// This response code is not used.
	ResponseCodeErrorUnableToGetUserOofSettings string = `ErrorUnableToGetUserOofSettings`
	// This error occurs when an unsuccessful attempt is made to remove an IM contact from a group. This error was introduced in Exchange 2013.
	ResponseCodeErrorUnableToRemoveImContactFromGroup string = `ErrorUnableToRemoveImContactFromGroup`
	// This error indicates that a user's dial plan is not available.
	ResponseCodeErrorUnifiedMessagingDialPlanNotFound string = `ErrorUnifiedMessagingDialPlanNotFound`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorUnifiedMessagingPromptNotFound string = `ErrorUnifiedMessagingPromptNotFound`
	// This error is intended for internal use only. This error was introduced in Exchange 2013.
	ResponseCodeErrorUnifiedMessagingReportDataNotFound string = `ErrorUnifiedMessagingReportDataNotFound`
	// This error indicates that the user could not be found.
	ResponseCodeErrorUnifiedMessagingRequestFailed string = `ErrorUnifiedMessagingRequestFailed`
	// This error indicates that a valid server for the dial plan can be found to handle the request.
	ResponseCodeErrorUnifiedMessagingServerNotFound string = `ErrorUnifiedMessagingServerNotFound`
	// This error occurs when you try to set the Culture property to a value that is not parsable by the System.Globalization.CultureInfo class.
	ResponseCodeErrorUnsupportedCulture string = `ErrorUnsupportedCulture`
	// This error occurs when a caller tries to use extended properties of types object, object array, error, or null.
	ResponseCodeErrorUnsupportedMapiPropertyType string = `ErrorUnsupportedMapiPropertyType`
	// This error occurs when you are trying to retrieve or set MIME content for an item other than a PostItem, Message, or CalendarItem object.
	ResponseCodeErrorUnsupportedMimeConversion string = `ErrorUnsupportedMimeConversion`
	// This error occurs when the caller passes a property that is invalid for a query. This can occur when calculated properties are used.
	ResponseCodeErrorUnsupportedPathForQuery string = `ErrorUnsupportedPathForQuery`
	// This error occurs when the caller passes a property that is invalid for a sort or group by property. This can occur when calculated properties are used.
	ResponseCodeErrorUnsupportedPathForSortGroup string = `ErrorUnsupportedPathForSortGroup`
	// This response code is not used.
	ResponseCodeErrorUnsupportedPropertyDefinition string = `ErrorUnsupportedPropertyDefinition`
	// This error indicates that the search folder restriction may be valid, but it is not supported by EWS.
	ResponseCodeErrorUnsupportedQueryFilter string = `ErrorUnsupportedQueryFilter`
	// This error indicates that the specified recurrence is not supported for tasks.
	ResponseCodeErrorUnsupportedRecurrence string = `ErrorUnsupportedRecurrence`
	// This response code is not used.
	ResponseCodeErrorUnsupportedSubFilter string = `ErrorUnsupportedSubFilter`
	// This error indicates that Exchange Web Services found a property type in the store but it cannot generate XML for the property type.
	ResponseCodeErrorUnsupportedTypeForConversion string = `ErrorUnsupportedTypeForConversion`
	// This error indicates that the delegate list failed to be saved after delegates were updated.
	ResponseCodeErrorUpdateDelegatesFailed string = `ErrorUpdateDelegatesFailed`
	// This error occurs when the single property path that is listed in a change description does not match the single property that is being set within the actual Item or Folder object.
	ResponseCodeErrorUpdatePropertyMismatch string = `ErrorUpdatePropertyMismatch`
	// This error indicates that the requester tried to grant permissions in its calendar or contacts folder to an external user but the sharing policy assigned to the requester indicates that the domain of the external user is not listed in the policy.
	ResponseCodeErrorUserNotAllowedByPolicy string = `ErrorUserNotAllowedByPolicy`
	// This error indicates that the requester is not enabled.
	ResponseCodeErrorUserNotUnifiedMessagingEnabled string = `ErrorUserNotUnifiedMessagingEnabled`
	// Indicates that the requester's organization has a set of federated domains but the requester's organization does not have any SMTP proxy addresses with one of the federated domains.
	ResponseCodeErrorUserWithoutFederatedProxyAddress string = `ErrorUserWithoutFederatedProxyAddress`
	// This error indicates that a calendar view start date or end date was set to 1/1/0001 12:00:00 AM or 12/31/9999 11:59:59 PM.
	ResponseCodeErrorValueOutOfRange string = `ErrorValueOutOfRange`
	// This error indicates that the Exchange store detected a virus in the message.
	ResponseCodeErrorVirusDetected string = `ErrorVirusDetected`
	// This error indicates that the Exchange store detected a virus in the message and deleted it.
	ResponseCodeErrorVirusMessageDeleted string = `ErrorVirusMessageDeleted`
	// This response code is not used.
	ResponseCodeErrorVoiceMailNotImplemented string = `ErrorVoiceMailNotImplemented`
	// This error is intended for internal use only.
	ResponseCodeErrorWeatherServiceDisabled string = `ErrorWeatherServiceDisabled`
	// This response code is not used.
	ResponseCodeErrorWebRequestInInvalidState string = `ErrorWebRequestInInvalidState`
	// This error indicates that there was an internal failure during communication with unmanaged code.
	ResponseCodeErrorWinTreeTwoInteropError string = `ErrorWin32InteropError`
	// This response code is not used.
	ResponseCodeErrorWorkingHoursSaveFailed string = `ErrorWorkingHoursSaveFailed`
	// This response code is not used.
	ResponseCodeErrorWorkingHoursXmlMalformed string = `ErrorWorkingHoursXmlMalformed`
	// This error indicates that a request can only be made to a server that is the same version as the mailbox server.
	ResponseCodeErrorWrongServerVersion string = `ErrorWrongServerVersion`
	// This error indicates that a request was made by a delegate that has a different server version than the principal's mailbox server.
	ResponseCodeErrorWrongServerVersionDelegate string = `ErrorWrongServerVersionDelegate`
	// No error occurred for the request.
	ResponseCodeNoError string = `NoError`
)

func (R *ResponseCode) SetForMarshal() {
	R.XMLName.Local = "m:ResponseCode"
}

func (R *ResponseCode) GetSchema() *Schema {
	return &SchemaMessages
}
