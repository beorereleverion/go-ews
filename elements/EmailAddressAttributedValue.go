package elements

// The EmailAddressAttributedValue element specifies an instance of an array of email addresses and their associated attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddressattributedvalue
type EmailAddressAttributedValue struct {
	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"t:Attributions"`
	// The Value element specifies the value of an EmailAddress associated with an attributions array.
	Value *ValueEmailAddressType `xml:"t:Value"`
}
