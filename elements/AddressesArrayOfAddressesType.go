package elements

// The Addresses element specifies an array of Address elements.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresses-arrayofaddressestype
type AddressesArrayOfAddressesType struct {
	// The Address element specifies the address of a contact.
	Address *AddressContactType `xml:"t:Address"`
}
