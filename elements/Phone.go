package elements

// The Phone element specifies a single phone number that results from a contact entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phone
type Phone struct {
	// The OriginalPhoneString element specifies the original phone number for a contact or persona.
	OriginalPhoneString *OriginalPhoneString `xml:"t:OriginalPhoneString"`
	// The PhoneString element specifies the phone number for an extracted contact.
	PhoneString *PhoneString `xml:"t:PhoneString"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"t:Type"`
}
