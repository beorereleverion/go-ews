package elements

// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/postaladdressattributedvalue
import "encoding/xml"

type PostalAddressAttributedValue struct {
	XMLName xml.Name

	// The Attributions element specifies an array of attributions for its associated Value element.
	Attributions *AttributionsArrayOfValueAttributionsType `xml:"Attributions"`
	// The Value element specifies information associated with a postal address.
	Value *ValuePersonaPostalAddressType `xml:"Value"`
}

func (P *PostalAddressAttributedValue) SetForMarshal() {
	P.XMLName.Local = "t:PostalAddressAttributedValue"
}

func (P *PostalAddressAttributedValue) GetSchema() *Schema {
	return &SchemaTypes
}
