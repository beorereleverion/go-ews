package elements

// The SendSMSAlertToRecipients element indicates the mobile phone numbers to which a Short Message Service (SMS) alert is to be sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendsmsalerttorecipients
type SendSMSAlertToRecipients struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
