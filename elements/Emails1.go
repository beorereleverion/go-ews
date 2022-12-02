package elements

// The Emails1 element specifies an array of EmailAddressAttributedValue values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emails1
type Emails1 struct {
	// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
	EmailAddressAttributedValue *EmailAddressAttributedValue `xml:"t:EmailAddressAttributedValue"`
}
