package elements

// The Phone element specifies a single phone number that results from a phone number entity extraction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phone-phoneentitytype
type PhonePhoneEntityType struct {
	// The OriginalPhoneString element specifies the original phone number for a contact or persona.
	OriginalPhoneString *OriginalPhoneString `xml:"t:OriginalPhoneString"`
	// The PhoneString element specifies the phone number for an extracted contact.
	PhoneString *PhoneString `xml:"t:PhoneString"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"t:Position"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"t:Type"`
}
