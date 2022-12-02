package elements

// The OtherAddresses element specifies an array of address values and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/otheraddresses
type OtherAddresses struct {
	// The PostalAddressAttributedValue element specifies an instance of an array of postal addresses and their associated attributions.
	PostalAddressAttributedValue *PostalAddressAttributedValue `xml:"t:PostalAddressAttributedValue"`
}
