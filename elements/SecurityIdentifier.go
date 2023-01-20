package elements

// The SecurityIdentifier element represents the security descriptor definition language (SDDL) form of a security identifier (SID).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/securityidentifier
import "encoding/xml"

type SecurityIdentifier struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SecurityIdentifier) SetForMarshal() {
	S.XMLName.Local = "t:SecurityIdentifier"
}

func (S *SecurityIdentifier) GetSchema() *Schema {
	return &SchemaTypes
}
