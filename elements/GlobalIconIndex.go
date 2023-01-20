package elements

// The GlobalIconIndex element identifies the global icon index for all items in a conversation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globaliconindex
import "encoding/xml"

type GlobalIconIndex struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// Specifies the appointment item icon.
	GlobalIconIndexAppointmentItem string = `AppointmentItem`
	// Specifies the meeting icon.
	GlobalIconIndexAppointmentMeet string = `AppointmentMeet`
	// Specifies the meeting cancel icon.
	GlobalIconIndexAppointmentMeetCancel string = `AppointmentMeetCancel`
	// Specifies the meeting information icon.
	GlobalIconIndexAppointmentMeetInfo string = `AppointmentMeetInfo`
	// Specifies the icon for a maybe response to the meeting.
	GlobalIconIndexAppointmentMeetMaybe string = `AppointmentMeetMaybe`
	// Specifies the icon for a tentative response to the meeting.
	GlobalIconIndexAppointmentMeetNY string = `AppointmentMeetNY`
	// Specifies the meeting declined icon.
	GlobalIconIndexAppointmentMeetNo string = `AppointmentMeetNo`
	// Specifies the recurring meeting icon.
	GlobalIconIndexAppointmentMeetRecur string = `AppointmentMeetRecur`
	// Specifies the meeting acceptance icon.
	GlobalIconIndexAppointmentMeetYes string = `AppointmentMeetYes`
	// Specifies the recurring appointment icon.
	GlobalIconIndexAppointmentRecur string = `AppointmentRecur`
	// Specifies the default icon.
	GlobalIconIndexDefault string = `Default`
	// Specifies the encrypted mail icon.
	GlobalIconIndexMailEncrypted string = `MailEncrypted`
	// Specifies the encrypted forwarded mail icon.
	GlobalIconIndexMailEncryptedForwarded string = `MailEncryptedForwarded`
	// Specifies the encrypted read mail icon.
	GlobalIconIndexMailEncryptedRead string = `MailEncryptedRead`
	// Specifies the encrypted replied to mail icon.
	GlobalIconIndexMailEncryptedReplied string = `MailEncryptedReplied`
	// Specifies the forwarded mail icon.
	GlobalIconIndexMailForwarded string = `MailForwarded`
	// Specifies the Information Rights Management (IRM)-protected mail icon.
	GlobalIconIndexMailIrm string = `MailIrm`
	// Specifies the IRM-protected forwarded mail icon.
	GlobalIconIndexMailIrmForwarded string = `MailIrmForwarded`
	// Specifies the IRM-protected replied to mail icon.
	GlobalIconIndexMailIrmReplied string = `MailIrmReplied`
	// Specifies the mail read icon.
	GlobalIconIndexMailRead string = `MailRead`
	// Specifies the replied to mail icon.
	GlobalIconIndexMailReplied string = `MailReplied`
	// Specifies the Secure/Multipurpose Internet Mail Extensions (S/MIME) signed mail icon.
	GlobalIconIndexMailSmimeSigned string = `MailSmimeSigned`
	// Specifies the S/MIME signed forwarded mail icon.
	GlobalIconIndexMailSmimeSignedForwarded string = `MailSmimeSignedForwarded`
	// Specifies the S/MIME signed read mail icon.
	GlobalIconIndexMailSmimeSignedRead string = `MailSmimeSignedRead`
	// Specifies the S/MIME signed replied to mail icon.
	GlobalIconIndexMailSmimeSignedReplied string = `MailSmimeSignedReplied`
	// Specifies the unread mail icon.
	GlobalIconIndexMailUnread string = `MailUnread`
	// Specifies the default icon for contacts.
	GlobalIconIndexOutlookDefaultForContacts string = `OutlookDefaultForContacts`
	// Specifies the icon for a post item.
	GlobalIconIndexPostItem string = `PostItem`
	// Specifies the SMS delivered mail icon.
	GlobalIconIndexSmsDelivered string = `SmsDelivered`
	// Specifies the icon for SMS routing to an external delivery point.
	GlobalIconIndexSmsRoutedToDeliveryPoint string = `SmsRoutedToDeliveryPoint`
	// Specifies the icon for SMS routing to an external messaging system.
	GlobalIconIndexSmsRoutedToExternalMessagingSystem string = `SmsRoutedToExternalMessagingSystem`
	// Specifies the icon for mail submitted for Short Message Service (SMS) routing.
	GlobalIconIndexSmsSubmitted string = `SmsSubmitted`
	// Specifies the task delegated icon.
	GlobalIconIndexTaskDelegated string = `TaskDelegated`
	// Specifies the task item icon.
	GlobalIconIndexTaskItem string = `TaskItem`
	// Specifies the task owned icon.
	GlobalIconIndexTaskOwned string = `TaskOwned`
	// Specifies the recurring task icon.
	GlobalIconIndexTaskRecur string = `TaskRecur`
)

func (G *GlobalIconIndex) SetForMarshal() {
	G.XMLName.Local = "t:GlobalIconIndex"
}

func (G *GlobalIconIndex) GetSchema() *Schema {
	return &SchemaTypes
}
