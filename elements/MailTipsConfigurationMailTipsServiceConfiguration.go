package elements

// The MailTipsConfiguration element contains service configuration information for the mail tips service.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailtipsconfiguration-mailtipsserviceconfiguration
type MailTipsConfigurationMailTipsServiceConfiguration struct {
	// The InternalDomains element identifies the list of internal SMTP domains of the organization.
	InternalDomains *InternalDomainsSmtpDomainList `xml:"t:InternalDomains"`
	// The LargeAudienceThreshold element represents the large audience threshold for a client.
	LargeAudienceThreshold *LargeAudienceThreshold `xml:"t:LargeAudienceThreshold"`
	// The MailTipsEnabled element indicates whether the mail tips service is available.
	MailTipsEnabled *MailTipsEnabled `xml:"t:MailTipsEnabled"`
	// The MaxMessageSize element represents the maximum message size a recipient can accept.
	MaxMessageSize *MaxMessageSize `xml:"t:MaxMessageSize"`
	// The MaxRecipientsPerGetMailTipsRequest element indicates the maximum number of recipients that can be passed to the GetMailTips operation.
	MaxRecipientsPerGetMailTipsRequest *MaxRecipientsPerGetMailTipsRequest `xml:"t:MaxRecipientsPerGetMailTipsRequest"`
	// The ShowExternalRecipientCount element indicates whether consumers of the GetMailTips operation have to show mail tips that indicate the number of external recipients to which a message is addressed.
	ShowExternalRecipientCount *ShowExternalRecipientCount `xml:"t:ShowExternalRecipientCount"`
}
