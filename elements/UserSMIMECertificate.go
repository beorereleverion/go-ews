package elements

// The UserSMIMECertificate element contains a value that encodes a contact's SMIME certificate.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/usersmimecertificate
import "encoding/xml"

type UserSMIMECertificate struct {
	XMLName xml.Name

	// The Base64Binary element contains a Base64-encoded value.
	Base64Binary *Base64Binary `xml:"Base64Binary"`
}

func (U *UserSMIMECertificate) SetForMarshal() {
	U.XMLName.Local = "t:UserSMIMECertificate"
}

func (U *UserSMIMECertificate) GetSchema() *Schema {
	return &SchemaTypes
}
