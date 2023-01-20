package elements

// The EmailAddress element defines the primary SMTP address of a mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddress-nonemptystringtype
import "encoding/xml"

type EmailAddressNonEmptyStringType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EmailAddressNonEmptyStringType) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddress"
}

func (E *EmailAddressNonEmptyStringType) GetSchema() *Schema {
	return &SchemaTypes
}
