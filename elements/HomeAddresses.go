package elements

// The HomeAddresses element specifies an array of home addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homeaddresses
type HomeAddresses struct {
	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"t:PostalAddressAttributedValue"`
}
