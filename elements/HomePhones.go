package elements

// The HomePhones element specifies an array of home phone numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homephones
type HomePhones struct {
	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"t:PhoneNumberAttributedValue"`
}
