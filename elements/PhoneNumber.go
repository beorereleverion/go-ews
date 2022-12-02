package elements

// The PhoneNumber element specifies the default phone number of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumber
type PhoneNumber struct {
	// The Number element specifies a phone number.
	Number *Number `xml:"t:Number"`
	// The Type element specifies the type of postal address or phone number, for example,HomeorBusiness.
	Type *Typestring `xml:"t:Type"`
}
