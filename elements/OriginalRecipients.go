package elements

// The OriginalRecipients element represents a list of e-mail addresses of the first message recipients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalrecipients
type OriginalRecipients struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
