package elements

// The IconIndex element identifies the icon index for an item or conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iconindex
type IconIndex string

const (
	// Specifies the appointment item icon.
	IconIndexAppointmentItem IconIndex = `AppointmentItem`
	// Specifies the meeting icon.
	IconIndexAppointmentMeet IconIndex = `AppointmentMeet`
	// Specifies the meeting cancel icon.
	IconIndexAppointmentMeetCancel IconIndex = `AppointmentMeetCancel`
	// Specifies the meeting information icon.
	IconIndexAppointmentMeetInfo IconIndex = `AppointmentMeetInfo`
	// Specifies the icon for a maybe response to the meeting.
	IconIndexAppointmentMeetMaybe IconIndex = `AppointmentMeetMaybe`
	// Specifies the icon for a tentative response to the meeting.
	IconIndexAppointmentMeetNY IconIndex = `AppointmentMeetNY`
	// Specifies the meeting declined icon.
	IconIndexAppointmentMeetNo IconIndex = `AppointmentMeetNo`
	// Specifies the recurring meeting icon.
	IconIndexAppointmentMeetRecur IconIndex = `AppointmentMeetRecur`
	// Specifies the meeting acceptance icon.
	IconIndexAppointmentMeetYes IconIndex = `AppointmentMeetYes`
	// Specifies the recurring appointment icon.
	IconIndexAppointmentRecur IconIndex = `AppointmentRecur`
	// Specifies the default icon.
	IconIndexDefault IconIndex = `Default`
	// Specifies the encrypted mail icon.
	IconIndexMailEncrypted IconIndex = `MailEncrypted`
	// Specifies the encrypted forwarded mail icon.
	IconIndexMailEncryptedForwarded IconIndex = `MailEncryptedForwarded`
	// Specifies the encrypted read mail icon.
	IconIndexMailEncryptedRead IconIndex = `MailEncryptedRead`
	// Specifies the encrypted replied to mail icon.
	IconIndexMailEncryptedReplied IconIndex = `MailEncryptedReplied`
	// Specifies the forwarded mail icon.
	IconIndexMailForwarded IconIndex = `MailForwarded`
	// Specifies the Information Rights Management (IRM)-protected mail icon.
	IconIndexMailIrm IconIndex = `MailIrm`
	// Specifies the IRM-protected forwarded mail icon.
	IconIndexMailIrmForwarded IconIndex = `MailIrmForwarded`
	// Specifies the IRM-protected replied to mail icon.
	IconIndexMailIrmReplied IconIndex = `MailIrmReplied`
	// Specifies the mail read icon.
	IconIndexMailRead IconIndex = `MailRead`
	// Specifies the replied to mail icon.
	IconIndexMailReplied IconIndex = `MailReplied`
	// Specifies the Secure/Multipurpose Internet Mail Extensions (S/MIME) signed mail icon.
	IconIndexMailSmimeSigned IconIndex = `MailSmimeSigned`
	// Specifies the S/MIME signed forwarded mail icon.
	IconIndexMailSmimeSignedForwarded IconIndex = `MailSmimeSignedForwarded`
	// Specifies the S/MIME signed read mail icon.
	IconIndexMailSmimeSignedRead IconIndex = `MailSmimeSignedRead`
	// Specifies the S/MIME signed replied to mail icon.
	IconIndexMailSmimeSignedReplied IconIndex = `MailSmimeSignedReplied`
	// Specifies the unread mail icon.
	IconIndexMailUnread IconIndex = `MailUnread`
	// Specifies the default icon for contacts.
	IconIndexOutlookDefaultForContacts IconIndex = `OutlookDefaultForContacts`
	// Specifies the icon for a post item.
	IconIndexPostItem IconIndex = `PostItem`
	// Specifies the SMS delivered mail icon.
	IconIndexSmsDelivered IconIndex = `SmsDelivered`
	// Specifies the icon for SMS routing to an external delivery point.
	IconIndexSmsRoutedToDeliveryPoint IconIndex = `SmsRoutedToDeliveryPoint`
	// Specifies the icon for SMS routing to an external messaging system.
	IconIndexSmsRoutedToExternalMessagingSystem IconIndex = `SmsRoutedToExternalMessagingSystem`
	// Specifies the icon mail submitted for Short Message Service (SMS) routing.
	IconIndexSmsSubmitted IconIndex = `SmsSubmitted`
	// Specifies the task delegated icon.
	IconIndexTaskDelegated IconIndex = `TaskDelegated`
	// Specifies the task item icon.
	IconIndexTaskItem IconIndex = `TaskItem`
	// Specifies the task owned icon.
	IconIndexTaskOwned IconIndex = `TaskOwned`
	// Specifies the recurring task icon.
	IconIndexTaskRecur IconIndex = `TaskRecur`
)
