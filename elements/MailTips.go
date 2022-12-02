package elements

// The MailTips element represents values for various types of mail tips.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtips
type MailTips struct {
	// The CustomMailTip element represents a customized mail tip message.
	CustomMailTip *CustomMailTip `xml:"t:CustomMailTip"`
	// The DeliveryRestricted element indicates whether delivery restrictions will prevent the sender's message from reaching the recipient.
	DeliveryRestricted *DeliveryRestricted `xml:"t:DeliveryRestricted"`
	// The ExternalMemberCount element represents the count of external members in a group.
	ExternalMemberCount *ExternalMemberCount `xml:"t:ExternalMemberCount"`
	// The InvalidRecipient element indicates whether the recipient is invalid.
	InvalidRecipient *InvalidRecipientMailTips `xml:"t:InvalidRecipient"`
	// The IsModerated element indicates whether the recipient's mailbox is being moderated.
	IsModerated *IsModerated `xml:"t:IsModerated"`
	// The MailboxFull element indicates whether the mailbox for the recipient is full.
	MailboxFull *MailboxFull `xml:"t:MailboxFull"`
	// The MaxMessageSize element represents the maximum message size a recipient can accept.
	MaxMessageSize *MaxMessageSize `xml:"t:MaxMessageSize"`
	// The OutOfOffice element represents the response message and a duration time for sending the response message.
	OutOfOffice *OutOfOffice `xml:"t:OutOfOffice"`
	// The PendingMailTips element indicates that the mail tips in this element could not be evaluated before the server's processing timeout expired.
	PendingMailTips *PendingMailTips `xml:"t:PendingMailTips"`
	// The RecipientAddress element represents the mailbox of the recipient.
	RecipientAddress *RecipientAddress `xml:"t:RecipientAddress"`
	// The TotalMemberCount element represents the count of all members in a group.
	TotalMemberCount *TotalMemberCount `xml:"t:TotalMemberCount"`
}
