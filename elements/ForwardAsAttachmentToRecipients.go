package elements

// The ForwardAsAttachmentToRecipients element indicates the e-mail addresses to which messages are to be forwarded as attachments.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardasattachmenttorecipients
type ForwardAsAttachmentToRecipients struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
