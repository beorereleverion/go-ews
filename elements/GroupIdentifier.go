package elements

// The GroupIdentifier element represents a single security identifier and attribute for an Active Directory directory service object group of which the account is a member.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupidentifier
import "encoding/xml"

type GroupIdentifier struct {
	XMLName xml.Name

	// The SID element represents the security descriptor definition language (SDDL) form of the security identifier (SID) for the account to use for impersonation or delegate access.
	SID *SID `xml:"SID"`
	// The SecurityIdentifier element represents the security descriptor definition language (SDDL) form of a security identifier (SID).
	SecurityIdentifier *SecurityIdentifier `xml:"SecurityIdentifier"`
	// Contains group attributes.
	Attributes *string `xml:"Attributes,attr"`
}

func (G *GroupIdentifier) SetForMarshal() {
	G.XMLName.Local = "t:GroupIdentifier"
}

func (G *GroupIdentifier) GetSchema() *Schema {
	return &SchemaTypes
}
