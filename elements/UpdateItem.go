package elements

// The UpdateItem element defines a request to update an item in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitem
type UpdateItem struct {
	// The ItemChange element contains an item identifier and the updates to apply to the item.
	ItemChange *ItemChange `xml:"t:ItemChange"`
	// The ItemChanges element contains an array of ItemChange elements that identify items and the updates to apply to the items.
	ItemChanges *ItemChanges `xml:"m:ItemChanges"`
	// The SavedItemFolderId element identifies the target folder for operations that update, send, and create items in a mailbox.
	SavedItemFolderId *SavedItemFolderId `xml:"m:SavedItemFolderId"`
	// Identifies the type of conflict resolution to try during an update. The default value is AutoResolve.
	ConflictResolution *string `xml:"ConflictResolution,attr"`
	// Describes how the item will be handled after it is updated. The MessageDisposition attribute is required for message items, including meeting messages such as meeting cancellations, meeting requests, and meeting responses.
	MessageDisposition *string `xml:"MessageDisposition,attr"`
	// Describes how meeting updates are communicated after a calendar item is updated. This attribute is required for calendar items and calendar item occurrences.
	SendMeetingInvitationsOrCancellations *string `xml:"SendMeetingInvitationsOrCancellations,attr"`
	// Indicates whether read receipts for the updated item should be suppressed. A text value of true indicates that read receipts should be suppressed. A value of false indicates that the read receipts will be sent to the sender. This attribute is optional.   This attribute was introduced in Exchange Server 2013 SP1.
	SuppressReadReceipts *string `xml:"SuppressReadReceipts,attr"`
}
