package elements

// The UserSMIMECertificate element contains a value that encodes a contact's SMIME certificate.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/usersmimecertificate
type UserSMIMECertificate struct {
	// The Base64Binary element contains a Base64-encoded value.
	Base64Binary *Base64Binary `xml:"t:Base64Binary"`
}
