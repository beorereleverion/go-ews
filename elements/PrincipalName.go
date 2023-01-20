package elements

// The PrincipalName element represents the user principal name (UPN) of the account to be used for Exchange impersonation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/principalname
import "encoding/xml"

type PrincipalName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PrincipalName) SetForMarshal() {
	P.XMLName.Local = "t:PrincipalName"
}

func (P *PrincipalName) GetSchema() *Schema {
	return &SchemaTypes
}
