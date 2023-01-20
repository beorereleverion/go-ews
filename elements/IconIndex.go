package elements

// The IconIndex element identifies the icon index for an item or conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/iconindex
import "encoding/xml"

type IconIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Specifies the appointment item icon.
	IconIndexAppointmentItem string = `AppointmentItem`
	// Specifies the meeting icon.
	IconIndexAppointmentMeet string = `AppointmentMeet`
	// Specifies the meeting cancel icon.
	IconIndexAppointmentMeetCancel string = `AppointmentMeetCancel`
	// Specifies the meeting information icon.
	IconIndexAppointmentMeetInfo string = `AppointmentMeetInfo`
	// Specifies the icon for a maybe response to the meeting.
	IconIndexAppointmentMeetMaybe string = `AppointmentMeetMaybe`
	// Specifies the icon for a tentative response to the meeting.
	IconIndexAppointmentMeetNY string = `AppointmentMeetNY`
	// Specifies the meeting declined icon.
	IconIndexAppointmentMeetNo string = `AppointmentMeetNo`
	// Specifies the recurring meeting icon.
	IconIndexAppointmentMeetRecur string = `AppointmentMeetRecur`
	// Specifies the meeting acceptance icon.
	IconIndexAppointmentMeetYes string = `AppointmentMeetYes`
	// Specifies the recurring appointment icon.
	IconIndexAppointmentRecur string = `AppointmentRecur`
	// Specifies the default icon.
	IconIndexDefault string = `Default`
	// Specifies the encrypted mail icon.
	IconIndexMailEncrypted string = `MailEncrypted`
	// Specifies the encrypted forwarded mail icon.
	IconIndexMailEncryptedForwarded string = `MailEncryptedForwarded`
	// Specifies the encrypted read mail icon.
	IconIndexMailEncryptedRead string = `MailEncryptedRead`
	// Specifies the encrypted replied to mail icon.
	IconIndexMailEncryptedReplied string = `MailEncryptedReplied`
	// Specifies the forwarded mail icon.
	IconIndexMailForwarded string = `MailForwarded`
	// Specifies the Information Rights Management (IRM)-protected mail icon.
	IconIndexMailIrm string = `MailIrm`
	// Specifies the IRM-protected forwarded mail icon.
	IconIndexMailIrmForwarded string = `MailIrmForwarded`
	// Specifies the IRM-protected replied to mail icon.
	IconIndexMailIrmReplied string = `MailIrmReplied`
	// Specifies the mail read icon.
	IconIndexMailRead string = `MailRead`
	// Specifies the replied to mail icon.
	IconIndexMailReplied string = `MailReplied`
	// Specifies the Secure/Multipurpose Internet Mail Extensions (S/MIME) signed mail icon.
	IconIndexMailSmimeSigned string = `MailSmimeSigned`
	// Specifies the S/MIME signed forwarded mail icon.
	IconIndexMailSmimeSignedForwarded string = `MailSmimeSignedForwarded`
	// Specifies the S/MIME signed read mail icon.
	IconIndexMailSmimeSignedRead string = `MailSmimeSignedRead`
	// Specifies the S/MIME signed replied to mail icon.
	IconIndexMailSmimeSignedReplied string = `MailSmimeSignedReplied`
	// Specifies the unread mail icon.
	IconIndexMailUnread string = `MailUnread`
	// Specifies the default icon for contacts.
	IconIndexOutlookDefaultForContacts string = `OutlookDefaultForContacts`
	// Specifies the icon for a post item.
	IconIndexPostItem string = `PostItem`
	// Specifies the SMS delivered mail icon.
	IconIndexSmsDelivered string = `SmsDelivered`
	// Specifies the icon for SMS routing to an external delivery point.
	IconIndexSmsRoutedToDeliveryPoint string = `SmsRoutedToDeliveryPoint`
	// Specifies the icon for SMS routing to an external messaging system.
	IconIndexSmsRoutedToExternalMessagingSystem string = `SmsRoutedToExternalMessagingSystem`
	// Specifies the icon mail submitted for Short Message Service (SMS) routing.
	IconIndexSmsSubmitted string = `SmsSubmitted`
	// Specifies the task delegated icon.
	IconIndexTaskDelegated string = `TaskDelegated`
	// Specifies the task item icon.
	IconIndexTaskItem string = `TaskItem`
	// Specifies the task owned icon.
	IconIndexTaskOwned string = `TaskOwned`
	// Specifies the recurring task icon.
	IconIndexTaskRecur string = `TaskRecur`
)

func (I *IconIndex) SetForMarshal() {
	I.XMLName.Local = "t:IconIndex"
}

func (I *IconIndex) GetSchema() *Schema {
	return &SchemaTypes
}
