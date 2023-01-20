package elements

// The PasswordExpirationDate element provides the password expiration date for a mailbox account.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/passwordexpirationdate
import "encoding/xml"

type PasswordExpirationDate struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PasswordExpirationDate) SetForMarshal() {
	P.XMLName.Local = "t:PasswordExpirationDate"
}

func (P *PasswordExpirationDate) GetSchema() *Schema {
	return &SchemaTypes
}
