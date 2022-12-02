package elements

// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumberattributedvalue
type PhoneNumberAttributedValue struct {
	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"t:Attributions"`
	// The Value element specifies a phone number and type information and is associated with a set of attributions.
	Value *ValuePersonaPhoneNumberType `xml:"t:Value"`
}
