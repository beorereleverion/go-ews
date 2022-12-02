package elements

// The HomePhones2 element specifies an array of HomePhone2 values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homephones2
type HomePhones2 struct {
	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"t:PhoneNumberAttributedValue"`
}
