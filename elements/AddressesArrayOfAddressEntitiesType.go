package elements

// The Addresses element specifies an array of AddressEntity elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresses-arrayofaddressentitiestype
import "encoding/xml"

type AddressesArrayOfAddressEntitiesType struct {
	XMLName xml.Name

	// The AddressEntity element specifies a single address entity.
	AddressEntity *AddressEntity `xml:"AddressEntity"`
}

func (A *AddressesArrayOfAddressEntitiesType) SetForMarshal() {
	A.XMLName.Local = "t:Addresses"
}

func (A *AddressesArrayOfAddressEntitiesType) GetSchema() *Schema {
	return &SchemaTypes
}
