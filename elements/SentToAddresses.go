package elements

// The SentToAddresses element indicates the e-mail addresses that incoming messages have to have been sent to in order for the condition or exception to apply.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senttoaddresses
type SentToAddresses struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
