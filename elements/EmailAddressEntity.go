package elements

// The EmailAddressEntity element specifies a single email address entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddressentity
type EmailAddressEntity struct {
	// The EmailAddress element specifies a single email address.
	EmailAddress *EmailAddressstring `xml:"t:EmailAddress"`
}
