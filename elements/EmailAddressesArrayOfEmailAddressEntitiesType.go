package elements

// The EmailAddresses element specifies an array of email address entities.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofemailaddressentitiestype
type EmailAddressesArrayOfEmailAddressEntitiesType struct {
	// The EmailAddressEntity element specifies a single email address entity.
	EmailAddressEntity *EmailAddressEntity `xml:"t:EmailAddressEntity"`
}
