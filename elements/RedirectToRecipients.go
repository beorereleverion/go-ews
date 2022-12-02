package elements

// The RedirectToRecipients element indicates the e-mail addresses to which messages are to be redirected.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/redirecttorecipients
type RedirectToRecipients struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
