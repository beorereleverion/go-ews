package elements

// The MSExchangeCertificate element contains a value that encodes the Microsoft Exchange certificate of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/msexchangecertificate
import "encoding/xml"

type MSExchangeCertificate struct {
	XMLName xml.Name

	// The Base64Binary element contains a Base64-encoded value.
	Base64Binary *Base64Binary `xml:"Base64Binary"`
}

func (M *MSExchangeCertificate) SetForMarshal() {
	M.XMLName.Local = "t:MSExchangeCertificate"
}

func (M *MSExchangeCertificate) GetSchema() *Schema {
	return &SchemaTypes
}
