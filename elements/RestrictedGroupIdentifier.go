package elements

// The RestrictGroupIdentifier element represents the group security identifier (SID) and attributes for a restricted group within a user token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/restrictedgroupidentifier
import "encoding/xml"

type RestrictedGroupIdentifier struct {
	XMLName xml.Name

	// The SecurityIdentifier element represents the security descriptor definition language (SDDL) form of a security identifier (SID).
	SecurityIdentifier *SecurityIdentifier `xml:"SecurityIdentifier"`
	// Contains group attributes.
	Attributes *string `xml:"Attributes,attr"`
}

func (R *RestrictedGroupIdentifier) SetForMarshal() {
	R.XMLName.Local = "t:RestrictedGroupIdentifier"
}

func (R *RestrictedGroupIdentifier) GetSchema() *Schema {
	return &SchemaTypes
}
