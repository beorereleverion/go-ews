package elements

// The EmailAddress element specifies a single email address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddress-string
import "encoding/xml"

type EmailAddressstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EmailAddressstring) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddress"
}

func (E *EmailAddressstring) GetSchema() *Schema {
	return &SchemaTypes
}
