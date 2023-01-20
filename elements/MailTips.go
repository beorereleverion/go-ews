package elements

// The MailTips element represents values for various types of mail tips.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtips
import "encoding/xml"

type MailTips struct {
	XMLName xml.Name

	// The CustomMailTip element represents a customized mail tip message.
	CustomMailTip *CustomMailTip `xml:"CustomMailTip"`
	// The DeliveryRestricted element indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	DeliveryRestricted *DeliveryRestricted `xml:"DeliveryRestricted"`
	// The ExternalMemberCount element represents the count of external members in a group.
	ExternalMemberCount *ExternalMemberCount `xml:"ExternalMemberCount"`
	// The InvalidRecipient element indicates whether the recipient is invalid.
	InvalidRecipient *InvalidRecipientMailTips `xml:"InvalidRecipient"`
	// The IsModerated element indicates whether the recipient's mailbox is being moderated.
	IsModerated *IsModerated `xml:"IsModerated"`
	// The MailboxFull element indicates whether the mailbox for the recipient is full.
	MailboxFull *MailboxFull `xml:"MailboxFull"`
	// The MaxMessageSize element represents the maximum message size a recipient can accept.
	MaxMessageSize *MaxMessageSize `xml:"MaxMessageSize"`
	// The OutOfOffice element represents the response message and a duration time for sending the response message.
	OutOfOffice *OutOfOffice `xml:"OutOfOffice"`
	// The PendingMailTips element indicates that the mail tips in this element could not be evaluated before the server's processing timeout expired.
	PendingMailTips *PendingMailTips `xml:"PendingMailTips"`
	// The RecipientAddress element represents the mailbox of the recipient.
	RecipientAddress *RecipientAddress `xml:"RecipientAddress"`
	// The TotalMemberCount element represents the count of all members in a group.
	TotalMemberCount *TotalMemberCount `xml:"TotalMemberCount"`
}

func (M *MailTips) SetForMarshal() {
	M.XMLName.Local = "m:MailTips"
}

func (M *MailTips) GetSchema() *Schema {
	return &SchemaMessages
}
