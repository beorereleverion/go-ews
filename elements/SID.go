package elements

// The SID element represents the security descriptor definition language (SDDL) form of the security identifier (SID) for the account to use for impersonation or delegate access.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sid
import "encoding/xml"

type SID struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SID) SetForMarshal() {
	S.XMLName.Local = "t:SID"
}

func (S *SID) GetSchema() *Schema {
	return &SchemaTypes
}
