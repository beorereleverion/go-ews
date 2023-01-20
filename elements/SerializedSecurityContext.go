package elements

// The SerializedSecurityContext element is used in the Simple Object Access Protocol (SOAP) header for token serialization in server-to-server authentication. Token serialization is not supported.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/serializedsecuritycontext
import "encoding/xml"

type SerializedSecurityContext struct {
	XMLName xml.Name

	// The GroupSids element represents a collection of Active Directory directory service group object security identifiers.
	GroupSids *GroupSids `xml:"GroupSids"`
	// The PrimarySmtpAddress element represents the primary Simple Mail Transfer Protocol (SMTP) address of an account to be used for server-to-server authorization or delegate access.
	PrimarySmtpAddress *PrimarySmtpAddress `xml:"PrimarySmtpAddress"`
	// The RestrictedGroupSids element represents a collection of restricted groups from a user's token.
	RestrictedGroupSids *RestrictedGroupSids `xml:"RestrictedGroupSids"`
	// The UserSid element represents the security descriptor definition language (SDDL) form of the user security identifier in a serialized security context SOAP header. Token serialization is not supported.
	UserSid *UserSid `xml:"UserSid"`
}

func (S *SerializedSecurityContext) SetForMarshal() {
	S.XMLName.Local = "t:SerializedSecurityContext"
}

func (S *SerializedSecurityContext) GetSchema() *Schema {
	return &SchemaTypes
}
