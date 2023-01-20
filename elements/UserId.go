package elements

// The UserId element identifies a delegate user or a user who has folder access permissions.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userid
import "encoding/xml"

type UserId struct {
	XMLName xml.Name

	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The DistinguishedUser element identifies Anonymous and Default user accounts for delegate access. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DistinguishedUser *DistinguishedUser `xml:"DistinguishedUser"`
	// The ExternalUserIdentity element identifies an external delegate user or an external user who has folder access permissions.
	ExternalUserIdentity *ExternalUserIdentity `xml:"ExternalUserIdentity"`
	// The PrimarySmtpAddress element represents the primary Simple Mail Transfer Protocol (SMTP) address of an account to be used for server-to-server authorization or delegate access.
	PrimarySmtpAddress *PrimarySmtpAddress `xml:"PrimarySmtpAddress"`
	// The SID element represents the security descriptor definition language (SDDL) form of the security identifier (SID) for the account to use for impersonation or delegate access.
	SID *SID `xml:"SID"`
}

func (U *UserId) SetForMarshal() {
	U.XMLName.Local = "t:UserId"
}

func (U *UserId) GetSchema() *Schema {
	return &SchemaTypes
}
