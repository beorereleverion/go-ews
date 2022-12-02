package elements

// The GlobalIconIndex element identifies the global icon index for all items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaliconindex
type GlobalIconIndex string

const (
	// Specifies the appointment item icon.
	GlobalIconIndexAppointmentItem GlobalIconIndex = `AppointmentItem`
	// Specifies the meeting icon.
	GlobalIconIndexAppointmentMeet GlobalIconIndex = `AppointmentMeet`
	// Specifies the meeting cancel icon.
	GlobalIconIndexAppointmentMeetCancel GlobalIconIndex = `AppointmentMeetCancel`
	// Specifies the meeting information icon.
	GlobalIconIndexAppointmentMeetInfo GlobalIconIndex = `AppointmentMeetInfo`
	// Specifies the icon for a maybe response to the meeting.
	GlobalIconIndexAppointmentMeetMaybe GlobalIconIndex = `AppointmentMeetMaybe`
	// Specifies the icon for a tentative response to the meeting.
	GlobalIconIndexAppointmentMeetNY GlobalIconIndex = `AppointmentMeetNY`
	// Specifies the meeting declined icon.
	GlobalIconIndexAppointmentMeetNo GlobalIconIndex = `AppointmentMeetNo`
	// Specifies the recurring meeting icon.
	GlobalIconIndexAppointmentMeetRecur GlobalIconIndex = `AppointmentMeetRecur`
	// Specifies the meeting acceptance icon.
	GlobalIconIndexAppointmentMeetYes GlobalIconIndex = `AppointmentMeetYes`
	// Specifies the recurring appointment icon.
	GlobalIconIndexAppointmentRecur GlobalIconIndex = `AppointmentRecur`
	// Specifies the default icon.
	GlobalIconIndexDefault GlobalIconIndex = `Default`
	// Specifies the encrypted mail icon.
	GlobalIconIndexMailEncrypted GlobalIconIndex = `MailEncrypted`
	// Specifies the encrypted forwarded mail icon.
	GlobalIconIndexMailEncryptedForwarded GlobalIconIndex = `MailEncryptedForwarded`
	// Specifies the encrypted read mail icon.
	GlobalIconIndexMailEncryptedRead GlobalIconIndex = `MailEncryptedRead`
	// Specifies the encrypted replied to mail icon.
	GlobalIconIndexMailEncryptedReplied GlobalIconIndex = `MailEncryptedReplied`
	// Specifies the forwarded mail icon.
	GlobalIconIndexMailForwarded GlobalIconIndex = `MailForwarded`
	// Specifies the Information Rights Management (IRM)-protected mail icon.
	GlobalIconIndexMailIrm GlobalIconIndex = `MailIrm`
	// Specifies the IRM-protected forwarded mail icon.
	GlobalIconIndexMailIrmForwarded GlobalIconIndex = `MailIrmForwarded`
	// Specifies the IRM-protected replied to mail icon.
	GlobalIconIndexMailIrmReplied GlobalIconIndex = `MailIrmReplied`
	// Specifies the mail read icon.
	GlobalIconIndexMailRead GlobalIconIndex = `MailRead`
	// Specifies the replied to mail icon.
	GlobalIconIndexMailReplied GlobalIconIndex = `MailReplied`
	// Specifies the Secure/Multipurpose Internet Mail Extensions (S/MIME) signed mail icon.
	GlobalIconIndexMailSmimeSigned GlobalIconIndex = `MailSmimeSigned`
	// Specifies the S/MIME signed forwarded mail icon.
	GlobalIconIndexMailSmimeSignedForwarded GlobalIconIndex = `MailSmimeSignedForwarded`
	// Specifies the S/MIME signed read mail icon.
	GlobalIconIndexMailSmimeSignedRead GlobalIconIndex = `MailSmimeSignedRead`
	// Specifies the S/MIME signed replied to mail icon.
	GlobalIconIndexMailSmimeSignedReplied GlobalIconIndex = `MailSmimeSignedReplied`
	// Specifies the unread mail icon.
	GlobalIconIndexMailUnread GlobalIconIndex = `MailUnread`
	// Specifies the default icon for contacts.
	GlobalIconIndexOutlookDefaultForContacts GlobalIconIndex = `OutlookDefaultForContacts`
	// Specifies the icon for a post item.
	GlobalIconIndexPostItem GlobalIconIndex = `PostItem`
	// Specifies the SMS delivered mail icon.
	GlobalIconIndexSmsDelivered GlobalIconIndex = `SmsDelivered`
	// Specifies the icon for SMS routing to an external delivery point.
	GlobalIconIndexSmsRoutedToDeliveryPoint GlobalIconIndex = `SmsRoutedToDeliveryPoint`
	// Specifies the icon for SMS routing to an external messaging system.
	GlobalIconIndexSmsRoutedToExternalMessagingSystem GlobalIconIndex = `SmsRoutedToExternalMessagingSystem`
	// Specifies the icon for mail submitted for Short Message Service (SMS) routing.
	GlobalIconIndexSmsSubmitted GlobalIconIndex = `SmsSubmitted`
	// Specifies the task delegated icon.
	GlobalIconIndexTaskDelegated GlobalIconIndex = `TaskDelegated`
	// Specifies the task item icon.
	GlobalIconIndexTaskItem GlobalIconIndex = `TaskItem`
	// Specifies the task owned icon.
	GlobalIconIndexTaskOwned GlobalIconIndex = `TaskOwned`
	// Specifies the recurring task icon.
	GlobalIconIndexTaskRecur GlobalIconIndex = `TaskRecur`
)
