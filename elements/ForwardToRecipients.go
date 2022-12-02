package elements

// The ForwardToRecipients element indicates the e-mail addresses to which messages are to be forwarded.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardtorecipients
type ForwardToRecipients struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
