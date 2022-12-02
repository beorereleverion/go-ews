package elements

// The GroupIdentifier element represents a single security identifier and attribute for an Active Directory directory service object group of which the account is a member.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupidentifier
type GroupIdentifier struct {
	// The SID element represents the security descriptor definition language (SDDL) form of the security identifier (SID) for the account to use for impersonation or delegate access.
	SID *SID `xml:"t:SID"`
	// The SecurityIdentifier element represents the security descriptor definition language (SDDL) form of a security identifier (SID).
	SecurityIdentifier *SecurityIdentifier `xml:"t:SecurityIdentifier"`
	// Contains group attributes.
	Attributes *string `xml:"Attributes,attr"`
}
