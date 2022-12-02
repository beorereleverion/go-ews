package elements

// The Value element specifies a phone number and type information and is associated with a set of attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-personaphonenumbertype
type ValuePersonaPhoneNumberType struct {
	// The Number element specifies a phone number.
	Number *Number `xml:"t:Number"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"t:Type"`
}
