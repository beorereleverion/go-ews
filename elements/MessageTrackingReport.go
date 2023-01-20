package elements

// The MessageTrackingReport element contains a single message that is returned in a GetMessageTrackingReport operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/messagetrackingreport
import "encoding/xml"

type MessageTrackingReport struct {
	XMLName xml.Name

	// The OriginalRecipients element represents a list of e-mail addresses of the first message recipients.
	OriginalRecipients *OriginalRecipients `xml:"OriginalRecipients"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The PurportedSender element contains contact information for the alleged sender of an e-mail message.
	PurportedSender *PurportedSender `xml:"PurportedSender"`
	// The RecipientTrackingEvents element represents a collection of one or more events for a message.
	RecipientTrackingEvents *RecipientTrackingEvents `xml:"RecipientTrackingEvents"`
	// The Sender element represents the e-mail address for the sender of the message.
	Sender *SenderEmailAddressType `xml:"Sender"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
	// The SubmitTime element represents the time at which the message was sent to the server.
	SubmitTime *SubmitTime `xml:"SubmitTime"`
}

func (M *MessageTrackingReport) SetForMarshal() {
	M.XMLName.Local = "m:MessageTrackingReport"
}

func (M *MessageTrackingReport) GetSchema() *Schema {
	return &SchemaMessages
}
