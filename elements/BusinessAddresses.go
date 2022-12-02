package elements

// The BusinessAddresses element specifies an array of business addresses and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/businessaddresses
type BusinessAddresses struct {
	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"t:PostalAddressAttributedValue"`
}
