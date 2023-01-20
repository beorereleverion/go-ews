package elements

// The HomeAddresses element specifies an array of home addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homeaddresses
import "encoding/xml"

type HomeAddresses struct {
	XMLName xml.Name

	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"PostalAddressAttributedValue"`
}

func (H *HomeAddresses) SetForMarshal() {
	H.XMLName.Local = "t:HomeAddresses"
}

func (H *HomeAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
