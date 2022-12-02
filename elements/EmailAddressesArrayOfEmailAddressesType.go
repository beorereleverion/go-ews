package elements

// The EmailAddresses element specifies an array of all email addresses of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofemailaddressestype
type EmailAddressesArrayOfEmailAddressesType struct {
	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"t:Address"`
}
