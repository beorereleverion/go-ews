package elements

// The MSExchangeCertificate element contains a value that encodes the Microsoft Exchange certificate of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/msexchangecertificate
type MSExchangeCertificate struct {
	// The Base64Binary element contains a Base64-encoded value.
	Base64Binary *Base64Binary `xml:"t:Base64Binary"`
}
