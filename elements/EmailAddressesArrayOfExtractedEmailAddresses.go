package elements

// The EmailAddresses element specifies an array of extracted email addresses.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofextractedemailaddresses
type EmailAddressesArrayOfExtractedEmailAddresses struct {
	// The EmailAddress element specifies a single email address.
	EmailAddress *EmailAddressstring `xml:"t:EmailAddress"`
}
