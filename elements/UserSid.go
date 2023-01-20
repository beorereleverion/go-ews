package elements

// The UserSid element represents the security descriptor definition language (SDDL) form of the user security identifier in a serialized security context SOAP header. Token serialization is not supported.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/usersid
import "encoding/xml"

type UserSid struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UserSid) SetForMarshal() {
	U.XMLName.Local = "t:UserSid"
}

func (U *UserSid) GetSchema() *Schema {
	return &SchemaTypes
}
