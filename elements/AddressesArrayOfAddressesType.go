package elements

// The Addresses element specifies an array of Address elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresses-arrayofaddressestype
import "encoding/xml"

type AddressesArrayOfAddressesType struct {
	XMLName xml.Name

	// The Address element specifies the address of a contact.
	Address *AddressContactType `xml:"Address"`
}

func (A *AddressesArrayOfAddressesType) SetForMarshal() {
	A.XMLName.Local = "t:Addresses"
}

func (A *AddressesArrayOfAddressesType) GetSchema() *Schema {
	return &SchemaTypes
}
