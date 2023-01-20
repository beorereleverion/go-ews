package elements

// The BusinessAddresses element specifies an array of business addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businessaddresses
import "encoding/xml"

type BusinessAddresses struct {
	XMLName xml.Name

	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"PostalAddressAttributedValue"`
}

func (B *BusinessAddresses) SetForMarshal() {
	B.XMLName.Local = "t:BusinessAddresses"
}

func (B *BusinessAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
