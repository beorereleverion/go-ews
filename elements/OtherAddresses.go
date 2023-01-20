package elements

// The OtherAddresses element specifies an array of address values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/otheraddresses
import "encoding/xml"

type OtherAddresses struct {
	XMLName xml.Name

	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"PostalAddressAttributedValue"`
}

func (O *OtherAddresses) SetForMarshal() {
	O.XMLName.Local = "t:OtherAddresses"
}

func (O *OtherAddresses) GetSchema() *Schema {
	return &SchemaTypes
}
