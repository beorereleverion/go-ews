package elements

// The EmailAddresses element specifies an array of all email addresses of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailaddresses-arrayofemailaddressestype
import "encoding/xml"

type EmailAddressesArrayOfEmailAddressesType struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (E *EmailAddressesArrayOfEmailAddressesType) SetForMarshal() {
	E.XMLName.Local = "t:EmailAddresses"
}

func (E *EmailAddressesArrayOfEmailAddressesType) GetSchema() *Schema {
	return &SchemaTypes
}
