package elements

// The Addresses element specifies an array of AddressEntity elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresses-arrayofaddressentitiestype
type AddressesArrayOfAddressEntitiesType struct {
	// The AddressEntity element specifies a single address entity.
	AddressEntity *AddressEntity `xml:"t:AddressEntity"`
}
