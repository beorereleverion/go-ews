package elements

// The EnhancedLocation element specifies location information such as the name, address, and optional notes about a location.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enhancedlocation
type EnhancedLocation struct {
	// The Annotation element contains optional notes added by a user.
	Annotation *Annotation `xml:"t:Annotation"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"t:DisplayName"`
	// The PostalAddress element specifies the postal address for a persona.
	PostalAddress *PostalAddressPersonaPostalAddressType `xml:"t:PostalAddress"`
}
