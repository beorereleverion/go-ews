package elements

// The EnhancedLocation element specifies location information such as the name, address, and optional notes about a location.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enhancedlocation
import "encoding/xml"

type EnhancedLocation struct {
	XMLName xml.Name

	// The Annotation element contains optional notes added by a user.
	Annotation *Annotation `xml:"Annotation"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The PostalAddress element specifies the postal address for a persona.
	PostalAddress *PostalAddressPersonaPostalAddressType `xml:"PostalAddress"`
}

func (E *EnhancedLocation) SetForMarshal() {
	E.XMLName.Local = "t:EnhancedLocation"
}

func (E *EnhancedLocation) GetSchema() *Schema {
	return &SchemaTypes
}
